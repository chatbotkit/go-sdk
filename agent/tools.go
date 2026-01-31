// Package agent provides default tools for agent execution.
//
// This file defines default tools that can be added to an agent for common
// file and shell operations. These tools mirror the functionality provided
// by the Node SDK's agent package.
package agent

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// DefaultTools returns a set of standard tools for file and shell operations.
// These tools provide:
//   - read: Read file contents with optional line ranges
//   - write: Write or modify file contents
//   - edit: Replace exact string occurrences in files
//   - exec: Execute shell commands with timeout
//
// Example:
//
//	tools := agent.DefaultTools()
//	events, errs := agent.ExecuteWithTools(ctx, client, agent.ExecuteWithToolsOptions{
//	    Model:    "gpt-4o",
//	    Messages: messages,
//	    Tools:    tools,
//	})
func DefaultTools() Tools {
	return Tools{
		"read": {
			Description: "Read the contents of a file. Supports optional line range to read specific sections.",
			Parameters: &Parameters{
				Properties: map[string]Property{
					"path": {
						Type:        "string",
						Description: "The file path to read",
					},
					"startLine": {
						Type:        "integer",
						Description: "The line number to start reading from (1-indexed)",
					},
					"endLine": {
						Type:        "integer",
						Description: "The line number to end reading at, inclusive (1-indexed)",
					},
				},
				Required: []string{"path"},
			},
			Handler: readFileHandler,
		},
		"write": {
			Description: "Write content to a file. Without line parameters, overwrites the entire file. With startLine only, inserts before that line. With startLine and endLine, replaces that range.",
			Parameters: &Parameters{
				Properties: map[string]Property{
					"path": {
						Type:        "string",
						Description: "The file path to write to",
					},
					"content": {
						Type:        "string",
						Description: "The content to write",
					},
					"startLine": {
						Type:        "integer",
						Description: "The line number to start writing at (1-indexed). If only startLine is provided, content is inserted before this line.",
					},
					"endLine": {
						Type:        "integer",
						Description: "The line number to end writing at, inclusive (1-indexed). Used with startLine to replace a range of lines.",
					},
				},
				Required: []string{"path", "content"},
			},
			Handler: writeFileHandler,
		},
		"edit": {
			Description: "Edit a file by replacing an exact string occurrence with a new string. Only one occurrence must exist.",
			Parameters: &Parameters{
				Properties: map[string]Property{
					"path": {
						Type:        "string",
						Description: "The file path to edit",
					},
					"oldString": {
						Type:        "string",
						Description: "The exact string to find and replace (must match exactly)",
					},
					"newString": {
						Type:        "string",
						Description: "The new string to replace with",
					},
				},
				Required: []string{"path", "oldString", "newString"},
			},
			Handler: editFileHandler,
		},
		"exec": {
			Description: "Execute a shell command (non-interactive only) using the current sh shell. Commands timeout after the specified duration (default 30 seconds). Use only for commands that run and exit automatically.",
			Parameters: &Parameters{
				Properties: map[string]Property{
					"command": {
						Type:        "string",
						Description: "The command to execute",
					},
					"timeout": {
						Type:        "number",
						Description: "Timeout in seconds. The command will be killed if it runs longer than this. Default is 30 seconds.",
					},
				},
				Required: []string{"command"},
			},
			Handler: execCommandHandler,
		},
	}
}

// readFileHandler handles the read tool.
func readFileHandler(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	path, ok := args["path"].(string)
	if !ok {
		return map[string]interface{}{
			"success": false,
			"error":   "path must be a string",
		}, nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}, nil
	}

	lines := strings.Split(string(content), "\n")
	totalLines := len(lines)

	// Check for line range parameters
	startLine, hasStart := getIntArg(args, "startLine")
	endLine, hasEnd := getIntArg(args, "endLine")

	// If no range specified, return full content
	if !hasStart && !hasEnd {
		return map[string]interface{}{
			"success":    true,
			"content":    string(content),
			"totalLines": totalLines,
		}, nil
	}

	// Convert 1-indexed to 0-indexed for slice operations
	start := 0
	if hasStart && startLine > 0 {
		start = startLine - 1
	}
	if start > totalLines {
		start = totalLines
	}

	end := totalLines
	if hasEnd && endLine > 0 {
		end = endLine
		if end > totalLines {
			end = totalLines
		}
	}

	outputContent := strings.Join(lines[start:end], "\n")

	result := map[string]interface{}{
		"success":    true,
		"content":    outputContent,
		"totalLines": totalLines,
	}
	if hasStart {
		result["startLine"] = startLine
	} else {
		result["startLine"] = 1
	}
	if hasEnd {
		result["endLine"] = endLine
	} else {
		result["endLine"] = totalLines
	}

	return result, nil
}

// writeFileHandler handles the write tool.
func writeFileHandler(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	path, ok := args["path"].(string)
	if !ok {
		return map[string]interface{}{
			"success": false,
			"error":   "path must be a string",
		}, nil
	}

	content, ok := args["content"].(string)
	if !ok {
		return map[string]interface{}{
			"success": false,
			"error":   "content must be a string",
		}, nil
	}

	startLine, hasStart := getIntArg(args, "startLine")
	endLine, hasEnd := getIntArg(args, "endLine")

	var finalContent string

	// Determine write mode based on parameters
	if !hasStart && !hasEnd {
		// Overwrite entire file
		finalContent = content
	} else {
		// Need to read existing content for line-based operations
		currentContent := ""
		existingData, err := os.ReadFile(path)
		if err == nil {
			currentContent = string(existingData)
		}
		// If file doesn't exist, treat as empty

		lines := strings.Split(currentContent, "\n")
		totalLines := len(lines)
		newLines := strings.Split(content, "\n")

		if hasStart && !hasEnd {
			// Insert before startLine
			insertIndex := startLine - 1
			if insertIndex < 0 {
				insertIndex = 0
			}
			if insertIndex > totalLines {
				insertIndex = totalLines
			}

			// Create new slice with inserted lines
			result := make([]string, 0, totalLines+len(newLines))
			result = append(result, lines[:insertIndex]...)
			result = append(result, newLines...)
			result = append(result, lines[insertIndex:]...)
			lines = result
		} else if hasStart && hasEnd {
			// Replace lines from startLine to endLine (inclusive)
			start := startLine - 1
			if start < 0 {
				start = 0
			}
			end := endLine
			if end > totalLines {
				end = totalLines
			}

			// Create new slice with replaced lines
			result := make([]string, 0, totalLines)
			result = append(result, lines[:start]...)
			result = append(result, newLines...)
			if end < totalLines {
				result = append(result, lines[end:]...)
			}
			lines = result
		} else {
			// endLine without startLine - treat as overwrite
			lines = newLines
		}

		finalContent = strings.Join(lines, "\n")
	}

	err := os.WriteFile(path, []byte(finalContent), 0644)
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}, nil
	}

	result := map[string]interface{}{
		"success": true,
	}
	if hasStart {
		result["startLine"] = startLine
	}
	if hasEnd {
		result["endLine"] = endLine
	}

	return result, nil
}

// editFileHandler handles the edit tool.
func editFileHandler(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	path, ok := args["path"].(string)
	if !ok {
		return map[string]interface{}{
			"success": false,
			"error":   "path must be a string",
		}, nil
	}

	oldString, ok := args["oldString"].(string)
	if !ok {
		return map[string]interface{}{
			"success": false,
			"error":   "oldString must be a string",
		}, nil
	}

	newString, ok := args["newString"].(string)
	if !ok {
		return map[string]interface{}{
			"success": false,
			"error":   "newString must be a string",
		}, nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}, nil
	}

	occurrences := strings.Count(string(content), oldString)

	if occurrences == 0 {
		return map[string]interface{}{
			"success": false,
			"error":   "String not found in file",
		}, nil
	}

	if occurrences > 1 {
		return map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Multiple occurrences found (%d). The old string must match exactly one location.", occurrences),
		}, nil
	}

	newContent := strings.Replace(string(content), oldString, newString, 1)

	err = os.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}, nil
	}

	// Truncate for preview
	oldPreview := oldString
	if len(oldPreview) > 100 {
		oldPreview = oldPreview[:100] + "..."
	}
	newPreview := newString
	if len(newPreview) > 100 {
		newPreview = newPreview[:100] + "..."
	}

	return map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Successfully replaced:\n  OLD: %s\n  NEW: %s", oldPreview, newPreview),
	}, nil
}

// execCommandHandler handles the exec tool.
func execCommandHandler(ctx context.Context, args map[string]interface{}) (interface{}, error) {
	command, ok := args["command"].(string)
	if !ok {
		return map[string]interface{}{
			"success": false,
			"error":   "command must be a string",
		}, nil
	}

	timeoutSecs := 30
	if t, ok := getIntArg(args, "timeout"); ok && t > 0 {
		timeoutSecs = t
	}

	timeout := time.Duration(timeoutSecs) * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sh", "-c", command)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		return map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Command timed out after %d seconds. This may indicate an interactive command.", timeoutSecs),
		}, nil
	}

	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			return map[string]interface{}{
				"success": false,
				"error":   fmt.Sprintf("Command exited with code %d", exitErr.ExitCode()),
				"stdout":  stdout.String(),
				"stderr":  stderr.String(),
			}, nil
		}
		return map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}, nil
	}

	return map[string]interface{}{
		"success": true,
		"stdout":  stdout.String(),
		"stderr":  stderr.String(),
	}, nil
}

// getIntArg extracts an integer argument from the args map.
// JSON numbers are typically float64 in Go's JSON unmarshaling.
func getIntArg(args map[string]interface{}, key string) (int, bool) {
	v, ok := args[key]
	if !ok {
		return 0, false
	}

	switch val := v.(type) {
	case float64:
		return int(val), true
	case int:
		return val, true
	case int64:
		return int(val), true
	default:
		return 0, false
	}
}
