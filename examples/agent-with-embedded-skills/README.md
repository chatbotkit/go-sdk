# Agent with Embedded Skills

Demonstrates bundling skill definitions directly into the executable using Go's `embed` package, so the binary ships self-contained with no external skill files.

## What it demonstrates

- Embedding a `skills/` directory at compile time with `//go:embed`
- Loading skills from an `embed.FS` via `agent.LoadSkillsFromFS`
- Using `fs.Sub` to re-root the embedded FS before passing it to the loader
- Passing the resulting skills feature to `agent.ExecuteWithTools` via `Extensions`

## Skills bundled in this example

| Skill         | Description                                                        |
| ------------- | ------------------------------------------------------------------ |
| `Code Review` | Reviews code snippets for correctness, style, and potential issues |
| `Summarize`   | Condenses long text into a concise summary                         |

Add more skills by creating a new subdirectory under `skills/` with a `SKILL.md` file - they are picked up automatically at compile time.

## Usage

```bash
export CHATBOTKIT_API_SECRET="your-api-key"

# Run with a custom task
go run . "Review this Go function: func add(a, b int) int { return a - b }"

# Run with the default demo task (summarize a passage)
go run .
```

## Example session

```
Loaded 2 embedded skill(s):
  • Code Review - Reviews code snippets for correctness, style, and potential issues...
  • Summarize - Condenses long text into a concise summary, preserving key points...

Task: Summarize the following text in bullet style: ...

--- Iteration 1 ---
• Go is a statically typed, compiled language designed at Google.
• Syntactically similar to C, with memory safety and garbage collection.
• Supports CSP-style concurrency and structural typing.
• Built for large-scale software engineering with fast compilation.

=== Agent exited with code 0 ===
```

## How it works

```go
//go:embed skills
var skillsFS embed.FS

// Re-root the FS so LoadSkillsFromFS sees skill directories at the top level.
subFS, _ := fs.Sub(skillsFS, "skills")
result, _ := agent.LoadSkillsFromFS(subFS)

feature := agent.CreateSkillsFeature(result.GetSkills())
```

### Skill directory layout

```
skills/
├── code-review/
│   └── SKILL.md   ← front matter: name + description
└── summarize/
    └── SKILL.md
```

Each `SKILL.md` must start with YAML front matter:

```markdown
---
name: My Skill
description: One-line description shown to the model.
---

# My Skill

Extended documentation...
```

## Difference from `agent.LoadSkills`

`LoadSkills` reads from OS directories at runtime and sets `SkillDefinition.Path` to an absolute filesystem path. `LoadSkillsFromFS` accepts any `fs.FS` - including `embed.FS` - and sets `Path` to the FS-relative directory name (e.g. `"code-review"`). The rest of the API is identical.
