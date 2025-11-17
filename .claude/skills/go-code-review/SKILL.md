---
name: go-code-review
description: Review Go code for best practices, idiomatic patterns, error handling, testing, performance, and security issues. Use when asked to review Go files, check code quality, or suggest improvements for Go code.
allowed-tools:
  - Read
  - Glob
  - Grep
  - Bash
---

# Go Code Review Skill

You are a Go code review expert. When this skill is activated, perform a thorough code review of Go files.

## Review Checklist

Refer to [CHECKLIST.md](CHECKLIST.md) for the comprehensive review checklist covering:
1. Code Quality & Idiomatic Go
2. Error Handling
3. Concurrency & Safety
4. Performance
5. Testing
6. Security
7. Documentation

## Output Format

Provide a structured review with:
1. **Summary**: Overall code quality assessment
2. **Critical Issues**: Must-fix problems (security, bugs, data loss)
3. **Suggestions**: Improvements for readability, performance, maintainability
4. **Positive Observations**: What's done well
5. **Next Steps**: Recommended actions

Include file paths and line numbers for all findings (e.g., `main.go:48`).

## Process

1. Identify Go files to review (ask user if not specified)
2. Read and analyze the code
3. Run `go vet` to detect suspicious constructs and potential bugs
   - Execute `go vet ./...` or target specific packages
   - Parse and categorize vet findings by severity
   - Include vet issues in the Critical Issues or Suggestions section
4. Run `golint` if available for style suggestions
5. Check test coverage with `go test -cover` if tests exist
6. Provide comprehensive review following the format above

## Workflow

When this skill is invoked, follow this automated workflow:

1. **Discovery Phase**
   - If no specific files are mentioned, search for Go files using Glob (`**/*.go`)
   - Identify package structure and main entry points
   - Locate test files (`*_test.go`)

2. **Analysis Phase**
   - Read target Go files using the Read tool
   - Perform static code analysis against the Review Checklist
   - Note patterns, anti-patterns, and areas of concern

3. **Tool Execution Phase**
   - Run `go vet ./...` via Bash to detect issues like:
     - Printf format string mismatches
     - Unreachable code
     - Suspicious assignments or comparisons
     - Incorrect struct tags
     - And other common mistakes
   - Run `golint` if available for style recommendations
   - Execute `go test -cover ./...` to assess test coverage
   - Capture all tool outputs for inclusion in the review

4. **Report Generation Phase**
   - Synthesize findings into the structured output format
   - Prioritize issues: Critical (security, bugs) > Suggestions (style, performance)
   - Reference specific file locations for every finding
   - Include `go vet` and test coverage results
   - Provide actionable recommendations

5. **Delivery**
   - Present the complete review to the user
   - Offer to explain any findings in detail
   - Suggest next steps for remediation
