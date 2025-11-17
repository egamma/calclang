# Go Code Review Checklist

## 1. Code Quality & Idiomatic Go
- Follow Go conventions (naming, formatting, package structure)
- Use gofmt/goimports formatting
- Proper use of Go idioms (e.g., accepting interfaces, returning structs)
- Avoid unnecessary complexity

## 2. Error Handling
- Check all error returns are handled
- Errors are wrapped with context when appropriate
- No silent error ignoring
- Proper error types and messages

## 3. Concurrency & Safety
- Proper goroutine management
- No race conditions
- Proper use of channels, mutexes, and sync primitives
- Check for potential deadlocks

## 4. Performance
- Unnecessary allocations
- Inefficient data structures
- Opportunities for optimization
- Proper use of pointers vs values

## 5. Testing
- Test coverage for critical paths
- Table-driven tests where appropriate
- Proper test naming conventions
- Edge cases covered

## 6. Security
- Input validation
- SQL injection risks
- Command injection risks
- Proper handling of sensitive data
- No hardcoded credentials

## 7. Documentation
- Exported functions have doc comments
- Package documentation exists
- Complex logic is commented
- Examples where helpful
