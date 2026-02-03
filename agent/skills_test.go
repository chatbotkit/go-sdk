package agent_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/chatbotkit/go-sdk/agent"
)

func TestParseFrontMatter(t *testing.T) {
	// Create a temporary directory with a SKILL.md file
	tmpDir := t.TempDir()
	skillDir := filepath.Join(tmpDir, "test-skill")
	if err := os.Mkdir(skillDir, 0755); err != nil {
		t.Fatalf("failed to create skill directory: %v", err)
	}

	skillContent := `---
name: Test Skill
description: A test skill for unit testing
---

# Test Skill

This is the skill documentation.
`

	skillPath := filepath.Join(skillDir, "SKILL.md")
	if err := os.WriteFile(skillPath, []byte(skillContent), 0644); err != nil {
		t.Fatalf("failed to write skill file: %v", err)
	}

	// Load skills from the temp directory
	result, err := agent.LoadSkills([]string{tmpDir})
	if err != nil {
		t.Fatalf("LoadSkills failed: %v", err)
	}

	if len(result.Skills) != 1 {
		t.Errorf("expected 1 skill, got %d", len(result.Skills))
	}

	skill := result.Skills[0]
	if skill.Name != "Test Skill" {
		t.Errorf("expected name 'Test Skill', got '%s'", skill.Name)
	}

	if skill.Description != "A test skill for unit testing" {
		t.Errorf("unexpected description: %s", skill.Description)
	}
}

func TestParseFrontMatterWithQuotes(t *testing.T) {
	tmpDir := t.TempDir()
	skillDir := filepath.Join(tmpDir, "quoted-skill")
	if err := os.Mkdir(skillDir, 0755); err != nil {
		t.Fatalf("failed to create skill directory: %v", err)
	}

	skillContent := `---
name: "Quoted Skill Name"
description: 'Quoted description'
---

# Quoted Skill
`

	skillPath := filepath.Join(skillDir, "SKILL.md")
	if err := os.WriteFile(skillPath, []byte(skillContent), 0644); err != nil {
		t.Fatalf("failed to write skill file: %v", err)
	}

	result, err := agent.LoadSkills([]string{tmpDir})
	if err != nil {
		t.Fatalf("LoadSkills failed: %v", err)
	}

	if len(result.Skills) != 1 {
		t.Errorf("expected 1 skill, got %d", len(result.Skills))
	}

	skill := result.Skills[0]
	if skill.Name != "Quoted Skill Name" {
		t.Errorf("expected name 'Quoted Skill Name', got '%s'", skill.Name)
	}

	if skill.Description != "Quoted description" {
		t.Errorf("unexpected description: %s", skill.Description)
	}
}

func TestLoadSkillsMissingFrontMatter(t *testing.T) {
	tmpDir := t.TempDir()
	skillDir := filepath.Join(tmpDir, "incomplete-skill")
	if err := os.Mkdir(skillDir, 0755); err != nil {
		t.Fatalf("failed to create skill directory: %v", err)
	}

	// Missing description
	skillContent := `---
name: Incomplete Skill
---

# Incomplete Skill
`

	skillPath := filepath.Join(skillDir, "SKILL.md")
	if err := os.WriteFile(skillPath, []byte(skillContent), 0644); err != nil {
		t.Fatalf("failed to write skill file: %v", err)
	}

	result, err := agent.LoadSkills([]string{tmpDir})
	if err != nil {
		t.Fatalf("LoadSkills failed: %v", err)
	}

	// Skill should be skipped due to missing description
	if len(result.Skills) != 0 {
		t.Errorf("expected 0 skills (invalid front matter), got %d", len(result.Skills))
	}
}

func TestLoadMultipleSkills(t *testing.T) {
	tmpDir := t.TempDir()

	// Create two skill directories
	for i, skillName := range []string{"skill-one", "skill-two"} {
		skillDir := filepath.Join(tmpDir, skillName)
		if err := os.Mkdir(skillDir, 0755); err != nil {
			t.Fatalf("failed to create skill directory: %v", err)
		}

		skillContent := "---\nname: " + skillName + "\ndescription: Description " + skillName + "\n---\n"
		skillPath := filepath.Join(skillDir, "SKILL.md")
		if err := os.WriteFile(skillPath, []byte(skillContent), 0644); err != nil {
			t.Fatalf("failed to write skill file: %v", err)
		}
		_ = i
	}

	result, err := agent.LoadSkills([]string{tmpDir})
	if err != nil {
		t.Fatalf("LoadSkills failed: %v", err)
	}

	if len(result.Skills) != 2 {
		t.Errorf("expected 2 skills, got %d", len(result.Skills))
	}
}

func TestCreateSkillsFeature(t *testing.T) {
	skills := []agent.SkillDefinition{
		{
			Name:        "Test Skill",
			Description: "A test skill",
			Path:        "/path/to/skill",
		},
	}

	feature := agent.CreateSkillsFeature(skills)

	if feature["name"] != "skills" {
		t.Errorf("expected feature name 'skills', got '%v'", feature["name"])
	}

	options, ok := feature["options"].(map[string]interface{})
	if !ok {
		t.Fatal("expected options to be a map")
	}

	skillsData, ok := options["skills"].([]map[string]string)
	if !ok {
		t.Fatal("expected skills to be a slice of maps")
	}

	if len(skillsData) != 1 {
		t.Errorf("expected 1 skill in feature, got %d", len(skillsData))
	}

	if skillsData[0]["name"] != "Test Skill" {
		t.Errorf("expected skill name 'Test Skill', got '%s'", skillsData[0]["name"])
	}
}

func TestGetSkills(t *testing.T) {
	tmpDir := t.TempDir()
	skillDir := filepath.Join(tmpDir, "test-skill")
	if err := os.Mkdir(skillDir, 0755); err != nil {
		t.Fatalf("failed to create skill directory: %v", err)
	}

	skillContent := "---\nname: Test Skill\ndescription: A test skill\n---\n"
	skillPath := filepath.Join(skillDir, "SKILL.md")
	if err := os.WriteFile(skillPath, []byte(skillContent), 0644); err != nil {
		t.Fatalf("failed to write skill file: %v", err)
	}

	result, err := agent.LoadSkills([]string{tmpDir})
	if err != nil {
		t.Fatalf("LoadSkills failed: %v", err)
	}

	skills := result.GetSkills()
	if len(skills) != 1 {
		t.Errorf("expected 1 skill from GetSkills, got %d", len(skills))
	}
}

func TestReload(t *testing.T) {
	tmpDir := t.TempDir()
	skillDir := filepath.Join(tmpDir, "test-skill")
	if err := os.Mkdir(skillDir, 0755); err != nil {
		t.Fatalf("failed to create skill directory: %v", err)
	}

	// Initially no SKILL.md
	result, err := agent.LoadSkills([]string{tmpDir})
	if err != nil {
		t.Fatalf("LoadSkills failed: %v", err)
	}

	if len(result.Skills) != 0 {
		t.Errorf("expected 0 skills initially, got %d", len(result.Skills))
	}

	// Add a SKILL.md file
	skillContent := "---\nname: New Skill\ndescription: A new skill\n---\n"
	skillPath := filepath.Join(skillDir, "SKILL.md")
	if err := os.WriteFile(skillPath, []byte(skillContent), 0644); err != nil {
		t.Fatalf("failed to write skill file: %v", err)
	}

	// Reload should pick up the new skill
	if err := result.Reload(); err != nil {
		t.Fatalf("Reload failed: %v", err)
	}

	if len(result.Skills) != 1 {
		t.Errorf("expected 1 skill after reload, got %d", len(result.Skills))
	}

	if result.Skills[0].Name != "New Skill" {
		t.Errorf("expected skill name 'New Skill', got '%s'", result.Skills[0].Name)
	}
}

func TestLoadSkillsNonExistentDirectory(t *testing.T) {
	result, err := agent.LoadSkills([]string{"/non/existent/directory"})
	if err != nil {
		t.Fatalf("LoadSkills should not fail for non-existent directory: %v", err)
	}

	if len(result.Skills) != 0 {
		t.Errorf("expected 0 skills for non-existent directory, got %d", len(result.Skills))
	}
}
