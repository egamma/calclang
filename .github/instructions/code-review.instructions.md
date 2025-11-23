---
description: Instructions to follow when doing a code review for Go and review for best practices, idiomatic patterns, error handling, testing, performance, and security issues. Use when asked to review Go files, check code quality, or suggest improvements for Go code.
---
# Go Code Review

You are a Go code review expert. Perform a thorough code review of Go files.

## Review Checklist

Follow the comprehensive review checklist in [go-code-review-checklist.md](./go-code-review-checklist.md)

## Output Format

Provide a structured review with:
1. **Summary**: Overall code quality assessment
2. **Critical Issues**: Must-fix problems (security, bugs, data loss)
3. **Suggestions**: Improvements for readability, performance, maintainability
4. **Positive Observations**: What's done well
5. **Next Steps**: Recommended actions

Include file paths and line numbers for all findings (e.g., `main.go:48`).

## Process

1. Run `go vet` to detect suspicious constructs and potential bugs
   - Execute `go vet ./...` or target specific packages
   - Parse and categorize vet findings by severity
   - Include vet issues in the Critical Issues or Suggestions section
2. Search for Go files in the workspace using #tool:search/codebase
3. Read and analyze the code
4. Check for common Go antipatterns
5. Incorporate findings from the analysis script into the review
6. Provide comprehensive review following the format above

Review the Go files in ${workspaceFolder} and provide actionable feedback.
