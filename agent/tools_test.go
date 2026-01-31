package agent_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/chatbotkit/go-sdk/agent"
)

func TestDefaultTools(t *testing.T) {
	tools := agent.DefaultTools()

	// Check that all expected tools are present
	expectedTools := []string{"read", "write", "edit", "exec"}
	for _, name := range expectedTools {
		if _, ok := tools[name]; !ok {
			t.Errorf("expected tool '%s' to be present", name)
		}
	}

	if len(tools) != len(expectedTools) {
		t.Errorf("expected %d tools, got %d", len(expectedTools), len(tools))
	}

	// Check that each tool has required fields
	for name, tool := range tools {
		if tool.Description == "" {
			t.Errorf("tool '%s' has empty description", name)
		}
		if tool.Handler == nil {
			t.Errorf("tool '%s' has nil handler", name)
		}
		if tool.Parameters == nil {
			t.Errorf("tool '%s' has nil parameters", name)
		}
	}
}

func TestReadTool(t *testing.T) {
	tools := agent.DefaultTools()
	readTool := tools["read"]
	ctx := context.Background()

	// Create a temporary file for testing
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	content := "line1\nline2\nline3\nline4\nline5"
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	t.Run("read full file", func(t *testing.T) {
		result, err := readTool.Handler(ctx, map[string]interface{}{
			"path": testFile,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if !m["success"].(bool) {
			t.Errorf("expected success=true, got false: %v", m["error"])
		}
		if m["content"].(string) != content {
			t.Errorf("unexpected content: %s", m["content"])
		}
		if m["totalLines"].(int) != 5 {
			t.Errorf("expected totalLines=5, got %d", m["totalLines"])
		}
	})

	t.Run("read with line range", func(t *testing.T) {
		result, err := readTool.Handler(ctx, map[string]interface{}{
			"path":      testFile,
			"startLine": 2.0, // JSON numbers are float64
			"endLine":   4.0,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if !m["success"].(bool) {
			t.Errorf("expected success=true: %v", m["error"])
		}
		expected := "line2\nline3\nline4"
		if m["content"].(string) != expected {
			t.Errorf("expected content=%q, got %q", expected, m["content"])
		}
	})

	t.Run("read non-existent file", func(t *testing.T) {
		result, err := readTool.Handler(ctx, map[string]interface{}{
			"path": "/non/existent/file.txt",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if m["success"].(bool) {
			t.Errorf("expected success=false for non-existent file")
		}
	})
}

func TestWriteTool(t *testing.T) {
	tools := agent.DefaultTools()
	writeTool := tools["write"]
	ctx := context.Background()

	t.Run("write new file", func(t *testing.T) {
		tmpDir := t.TempDir()
		testFile := filepath.Join(tmpDir, "new.txt")
		content := "hello world"

		result, err := writeTool.Handler(ctx, map[string]interface{}{
			"path":    testFile,
			"content": content,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if !m["success"].(bool) {
			t.Errorf("expected success=true: %v", m["error"])
		}

		// Verify content
		data, err := os.ReadFile(testFile)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(data) != content {
			t.Errorf("expected content=%q, got %q", content, string(data))
		}
	})

	t.Run("insert at line", func(t *testing.T) {
		tmpDir := t.TempDir()
		testFile := filepath.Join(tmpDir, "insert.txt")
		if err := os.WriteFile(testFile, []byte("line1\nline2\nline3"), 0644); err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}

		result, err := writeTool.Handler(ctx, map[string]interface{}{
			"path":      testFile,
			"content":   "inserted",
			"startLine": 2.0,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if !m["success"].(bool) {
			t.Errorf("expected success=true: %v", m["error"])
		}

		// Verify content
		data, err := os.ReadFile(testFile)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		expected := "line1\ninserted\nline2\nline3"
		if string(data) != expected {
			t.Errorf("expected content=%q, got %q", expected, string(data))
		}
	})

	t.Run("replace line range", func(t *testing.T) {
		tmpDir := t.TempDir()
		testFile := filepath.Join(tmpDir, "replace.txt")
		if err := os.WriteFile(testFile, []byte("line1\nline2\nline3\nline4"), 0644); err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}

		result, err := writeTool.Handler(ctx, map[string]interface{}{
			"path":      testFile,
			"content":   "replaced",
			"startLine": 2.0,
			"endLine":   3.0,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if !m["success"].(bool) {
			t.Errorf("expected success=true: %v", m["error"])
		}

		// Verify content
		data, err := os.ReadFile(testFile)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		expected := "line1\nreplaced\nline4"
		if string(data) != expected {
			t.Errorf("expected content=%q, got %q", expected, string(data))
		}
	})
}

func TestEditTool(t *testing.T) {
	tools := agent.DefaultTools()
	editTool := tools["edit"]
	ctx := context.Background()

	t.Run("replace single occurrence", func(t *testing.T) {
		tmpDir := t.TempDir()
		testFile := filepath.Join(tmpDir, "edit.txt")
		if err := os.WriteFile(testFile, []byte("hello world"), 0644); err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}

		result, err := editTool.Handler(ctx, map[string]interface{}{
			"path":      testFile,
			"oldString": "world",
			"newString": "universe",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if !m["success"].(bool) {
			t.Errorf("expected success=true: %v", m["error"])
		}

		// Verify content
		data, err := os.ReadFile(testFile)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}
		if string(data) != "hello universe" {
			t.Errorf("expected content='hello universe', got %q", string(data))
		}
	})

	t.Run("fail on multiple occurrences", func(t *testing.T) {
		tmpDir := t.TempDir()
		testFile := filepath.Join(tmpDir, "multi.txt")
		if err := os.WriteFile(testFile, []byte("foo bar foo"), 0644); err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}

		result, err := editTool.Handler(ctx, map[string]interface{}{
			"path":      testFile,
			"oldString": "foo",
			"newString": "baz",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if m["success"].(bool) {
			t.Errorf("expected success=false for multiple occurrences")
		}
		if !strings.Contains(m["error"].(string), "Multiple occurrences") {
			t.Errorf("expected error about multiple occurrences, got: %v", m["error"])
		}
	})

	t.Run("fail on no match", func(t *testing.T) {
		tmpDir := t.TempDir()
		testFile := filepath.Join(tmpDir, "nomatch.txt")
		if err := os.WriteFile(testFile, []byte("hello world"), 0644); err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}

		result, err := editTool.Handler(ctx, map[string]interface{}{
			"path":      testFile,
			"oldString": "notfound",
			"newString": "replacement",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if m["success"].(bool) {
			t.Errorf("expected success=false for no match")
		}
	})
}

func TestExecTool(t *testing.T) {
	tools := agent.DefaultTools()
	execTool := tools["exec"]
	ctx := context.Background()

	t.Run("successful command", func(t *testing.T) {
		result, err := execTool.Handler(ctx, map[string]interface{}{
			"command": "echo hello",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if !m["success"].(bool) {
			t.Errorf("expected success=true: %v", m["error"])
		}
		if !strings.Contains(m["stdout"].(string), "hello") {
			t.Errorf("expected stdout to contain 'hello', got: %q", m["stdout"])
		}
	})

	t.Run("failing command", func(t *testing.T) {
		result, err := execTool.Handler(ctx, map[string]interface{}{
			"command": "exit 1",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if m["success"].(bool) {
			t.Errorf("expected success=false for failing command")
		}
	})

	t.Run("command with stderr", func(t *testing.T) {
		result, err := execTool.Handler(ctx, map[string]interface{}{
			"command": "echo error >&2",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if !m["success"].(bool) {
			t.Errorf("expected success=true: %v", m["error"])
		}
		if !strings.Contains(m["stderr"].(string), "error") {
			t.Errorf("expected stderr to contain 'error', got: %q", m["stderr"])
		}
	})

	t.Run("command with custom timeout", func(t *testing.T) {
		result, err := execTool.Handler(ctx, map[string]interface{}{
			"command": "echo quick",
			"timeout": 5.0,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		m := result.(map[string]interface{})
		if !m["success"].(bool) {
			t.Errorf("expected success=true: %v", m["error"])
		}
	})
}
