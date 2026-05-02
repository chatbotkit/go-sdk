// Package agent provides skills loading functionality for ChatBotKit.
//
// This file provides utilities for loading skill definitions from local
// directories or any fs.FS (including embed.FS) that can be passed to the
// ChatBotKit API as a feature.
package agent

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"gopkg.in/yaml.v3"
)

// SkillDefinition represents a skill definition loaded from a SKILL.md file.
type SkillDefinition struct {
	// Name is the skill name from front matter.
	Name string
	// Description is the skill description from front matter.
	Description string
	// Path is the location of the skill directory. When loaded from the OS
	// filesystem this is an absolute path; when loaded from an fs.FS it is
	// the FS-relative path (e.g. "hello-world").
	Path string
}

// SkillsResult contains loaded skills and a reload function.
type SkillsResult struct {
	// Skills is the list of loaded skill definitions.
	Skills []SkillDefinition
	// mu protects the Skills slice.
	mu sync.RWMutex
	// directories holds OS paths used by LoadSkills.
	directories []string
	// fsys holds the FS used by LoadSkillsFromFS.
	fsys fs.FS
}

// GetSkills returns a copy of the current skills.
func (r *SkillsResult) GetSkills() []SkillDefinition {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]SkillDefinition, len(r.Skills))
	copy(result, r.Skills)
	return result
}

// Reload rescans the source for skill definitions.
// For OS-based results this re-reads the directories; for fs.FS-based results
// it rescans the embedded filesystem (useful when the FS is backed by an
// os.DirFS or similar mutable source).
func (r *SkillsResult) Reload() error {
	var newSkills []SkillDefinition

	if r.fsys != nil {
		skills, err := scanFS(r.fsys, "")
		if err != nil {
			return err
		}
		newSkills = skills
	} else {
		for _, dir := range r.directories {
			absDir, err := filepath.Abs(dir)
			if err != nil {
				continue
			}
			skills, err := scanFS(os.DirFS(absDir), absDir)
			if err != nil {
				continue
			}
			newSkills = append(newSkills, skills...)
		}
	}

	r.mu.Lock()
	r.Skills = newSkills
	r.mu.Unlock()

	return nil
}

// frontMatter represents the YAML front matter structure.
type frontMatter struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

// frontMatterRegex matches YAML front matter in markdown files.
var frontMatterRegex = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---`)

// parseFrontMatter parses YAML front matter from markdown content.
func parseFrontMatter(content string) (name, description string) {
	match := frontMatterRegex.FindStringSubmatch(content)
	if match == nil || len(match) < 2 {
		return "", ""
	}

	var fm frontMatter
	if err := yaml.Unmarshal([]byte(match[1]), &fm); err != nil {
		return "", ""
	}

	return fm.Name, fm.Description
}

// loadSkillFromFS loads a single skill from a named subdirectory within fsys.
// basePath is prepended to dirName when constructing SkillDefinition.Path;
// pass an empty string for FS-relative paths (e.g. embedded skills).
func loadSkillFromFS(fsys fs.FS, dirName string, basePath string) (*SkillDefinition, error) {
	data, err := fs.ReadFile(fsys, dirName+"/SKILL.md")
	if err != nil {
		return nil, err
	}

	name, description := parseFrontMatter(string(data))
	if name == "" || description == "" {
		return nil, nil
	}

	path := dirName
	if basePath != "" {
		path = filepath.Join(basePath, dirName)
	}

	return &SkillDefinition{
		Name:        name,
		Description: description,
		Path:        path,
	}, nil
}

// scanFS scans the root of fsys for skill subdirectories.
// basePath is passed through to loadSkillFromFS for path construction.
func scanFS(fsys fs.FS, basePath string) ([]SkillDefinition, error) {
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var skills []SkillDefinition
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		skill, err := loadSkillFromFS(fsys, entry.Name(), basePath)
		if err != nil || skill == nil {
			continue
		}

		skills = append(skills, *skill)
	}

	return skills, nil
}

// LoadSkills loads skills from multiple OS directories.
// Each directory should contain subdirectories with SKILL.md files.
// Call Reload() to rescan directories for changes.
func LoadSkills(directories []string) (*SkillsResult, error) {
	result := &SkillsResult{
		Skills:      make([]SkillDefinition, 0),
		directories: directories,
	}

	for _, dir := range directories {
		absDir, err := filepath.Abs(dir)
		if err != nil {
			// @note continue on error to match prior behavior
			continue
		}

		skills, err := scanFS(os.DirFS(absDir), absDir)
		if err != nil {
			continue
		}
		result.Skills = append(result.Skills, skills...)
	}

	return result, nil
}

// LoadSkillsFromFS loads skills from any fs.FS, including embed.FS.
// The FS root is scanned for skill subdirectories, each containing a SKILL.md.
// SkillDefinition.Path will be the FS-relative directory name (e.g. "hello-world").
//
// Typical usage with Go embedding:
//
//	//go:embed skills
//	var skillsFS embed.FS
//
//	subFS, _ := fs.Sub(skillsFS, "skills")
//	result, err := agent.LoadSkillsFromFS(subFS)
func LoadSkillsFromFS(fsys fs.FS) (*SkillsResult, error) {
	result := &SkillsResult{
		Skills: make([]SkillDefinition, 0),
		fsys:   fsys,
	}

	skills, err := scanFS(fsys, "")
	if err != nil {
		return result, nil
	}
	result.Skills = skills

	return result, nil
}

// CreateSkillsFeature creates a skills feature configuration from skill definitions.
// This can be passed to the extensions.features array when calling the API.
func CreateSkillsFeature(skills []SkillDefinition) map[string]interface{} {
	skillsData := make([]map[string]string, len(skills))
	for i, skill := range skills {
		skillsData[i] = map[string]string{
			"name":        skill.Name,
			"description": skill.Description,
			"path":        skill.Path,
		}
	}

	return map[string]interface{}{
		"name": "skills",
		"options": map[string]interface{}{
			"skills": skillsData,
		},
	}
}
