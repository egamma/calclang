---
argument-hint: the github issue number
description: Fix a github issue using TDD
---

Please analyze and fix the GitHub issue: $ARGUMENTS.

Follow these steps:

1. Use `gh issue view` to get the issue details
2. Use Test-Driven Development (TDD) to fix the issue
3. Ensure code passes linting and type checking

Remember to use the GitHub CLI (`gh`) for all GitHub-related tasks.

#### Test-Driven Development

Use Test-Driven Development:

1. Write failing unit test first (Jest).
   IMPORTANT: Always start by writing a unit test (not an integration or E2E test).
2. Write minimal code to pass test
3. Refactor while keeping tests green
