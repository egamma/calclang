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

You are a Go code review expert. When this skill is activated, perform a thorough code review of Go files focusing on:

## Review Checklist

### 1. Code Quality & Idiomatic Go
- Follow Go conventions (naming, formatting, package structure)
- Use gofmt/goimports formatting
- Proper use of Go idioms (e.g., accepting interfaces, returning structs)
- Avoid unnecessary complexity

### 2. Error Handling
- Check all error returns are handled
- Errors are wrapped with context when appropriate
- No silent error ignoring
- Proper error types and messages

### 3. Concurrency & Safety
- Proper goroutine management
- No race conditions
- Proper use of channels, mutexes, and sync primitives
- Check for potential deadlocks

### 4. Performance
- Unnecessary allocations
- Inefficient data structures
- Opportunities for optimization
- Proper use of pointers vs values

### 5. Testing
- Test coverage for critical paths
- Table-driven tests where appropriate
- Proper test naming conventions
- Edge cases covered

### 6. Security
- Input validation
- SQL injection risks
- Command injection risks
- Proper handling of sensitive data
- No hardcoded credentials

### 7. Documentation
- Exported functions have doc comments
- Package documentation exists
- Complex logic is commented
- Examples where helpful

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
3. Run `go vet` and `golint` if available
4. Check test coverage with `go test -cover` if tests exist
5. Provide comprehensive review following the format above
