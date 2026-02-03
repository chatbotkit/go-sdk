// Package agent provides skills loading functionality for ChatBotKit.
//
// This file provides utilities for loading skill definitions from local directories
// that can be passed to the ChatBotKit API as a feature.
package agent

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

// SkillDefinition represents a skill definition loaded from a SKILL.md file.
type SkillDefinition struct {
	// Name is the skill name from front matter.
	Name string
	// Description is the skill description from front matter.
	Description string
	// Path is the absolute path to the skill directory.
	Path string
}

// SkillsResult contains loaded skills and a reload function.
type SkillsResult struct {
	// Skills is the list of loaded skill definitions.
	Skills []SkillDefinition
	// mu protects the Skills slice.
	mu sync.RWMutex
	// directories to scan
	directories []string
}

// GetSkills returns a copy of the current skills.
func (r *SkillsResult) GetSkills() []SkillDefinition {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]SkillDefinition, len(r.Skills))
	copy(result, r.Skills)
	return result
}

// Reload rescans the directories for skill definitions.
func (r *SkillsResult) Reload() error {
	newSkills := make([]SkillDefinition, 0)

	for _, dir := range r.directories {
		skills, err := scanDirectory(dir)
		if err != nil {
			continue
		}
		newSkills = append(newSkills, skills...)
	}

	r.mu.Lock()
	r.Skills = newSkills
	r.mu.Unlock()

	return nil
}

// frontMatter represents the YAML front matter structure
type frontMatter struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

// frontMatterRegex matches YAML front matter in markdown files
var frontMatterRegex = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---`)

// parseFrontMatter parses YAML front matter from markdown content.
// Expects format:
// ---
// name: Skill Name
// description: Skill description
// ---
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

// loadSkillFromDirectory loads a single skill from a directory containing a SKILL.md file.
func loadSkillFromDirectory(skillDir string) (*SkillDefinition, error) {
	skillFilePath := filepath.Join(skillDir, "SKILL.md")

	file, err := os.Open(skillFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read file content
	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text())
		content.WriteString("\n")
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	name, description := parseFrontMatter(content.String())
	if name == "" || description == "" {
		return nil, nil
	}

	absPath, err := filepath.Abs(skillDir)
	if err != nil {
		return nil, err
	}

	return &SkillDefinition{
		Name:        name,
		Description: description,
		Path:        absPath,
	}, nil
}

// scanDirectory scans a directory for skill subdirectories.
func scanDirectory(baseDir string) ([]SkillDefinition, error) {
	entries, err := os.ReadDir(baseDir)
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

		entryPath := filepath.Join(baseDir, entry.Name())
		skill, err := loadSkillFromDirectory(entryPath)
		if err != nil || skill == nil {
			continue
		}

		skills = append(skills, *skill)
	}

	return skills, nil
}

// LoadSkills loads skills from multiple directories.
// Each directory should contain subdirectories with SKILL.md files.
// Call Reload() to rescan directories for changes.
func LoadSkills(directories []string) (*SkillsResult, error) {
	result := &SkillsResult{
		Skills:      make([]SkillDefinition, 0),
		directories: directories,
	}

	for _, dir := range directories {
		skills, err := scanDirectory(dir)
		if err != nil {
			// @note continue on error to match Node SDK behavior
			continue
		}
		result.Skills = append(result.Skills, skills...)
	}

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
