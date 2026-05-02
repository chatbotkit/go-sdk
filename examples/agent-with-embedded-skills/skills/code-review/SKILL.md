---
name: Code Review
description: Reviews code snippets for correctness, style, and potential issues. Provides actionable feedback.
---

# Code Review Skill

Analyzes provided code snippets and returns structured feedback covering:

- **Correctness**: logic errors, edge cases, off-by-one errors
- **Style**: naming conventions, readability, idiomatic patterns
- **Security**: common vulnerabilities, unsafe patterns
- **Performance**: inefficient algorithms, unnecessary allocations

## Usage

Pass a code snippet to this skill along with the language name. The skill returns a list of findings, each with a severity (info, warning, error) and a suggested fix.
