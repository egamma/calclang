#!/bin/bash

# Go Code Analysis Script
# Runs static analysis and test coverage checks

set -e  # Exit on error

echo "=== Running Go Static Analysis ==="
echo

echo "Running go vet..."
if go vet ./...; then
    echo "✓ go vet: No issues found"
else
    echo "✗ go vet: Issues detected (see above)"
fi

echo
echo "Running tests with coverage..."
if go test -cover ./...; then
    echo "✓ Tests completed"
else
    echo "✗ Tests failed or not found"
fi

echo
echo "=== Analysis Complete ==="
