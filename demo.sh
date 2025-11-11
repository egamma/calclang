#!/bin/bash

# CALCLANG Demo Script
# Demonstrates all arithmetic operations: ADD, SUB, MUL, DIV

echo "CALCLANG Demo - Basic Arithmetic Operations"
echo "============================================"
echo ""

# Build the program
echo "Building CALCLANG..."
go build -o calclang main.go
echo ""

# Demo addition
echo "1. Addition: ADD 10 5"
echo "   Result: $(echo "ADD 10 5" | ./calclang | grep '^>' | head -n 1 | sed 's/^> //')"
echo ""

# Demo subtraction
echo "2. Subtraction: SUB 10 5"
echo "   Result: $(echo "SUB 10 5" | ./calclang | grep '^>' | head -n 1 | sed 's/^> //')"
echo ""

# Demo multiplication
echo "3. Multiplication: MUL 10 5"
echo "   Result: $(echo "MUL 10 5" | ./calclang | grep '^>' | head -n 1 | sed 's/^> //')"
echo ""

# Demo division
echo "4. Division: DIV 10 5"
echo "   Result: $(echo "DIV 10 5" | ./calclang | grep '^>' | head -n 1 | sed 's/^> //')"
echo ""

# Demo division by zero
echo "5. Division by zero: DIV 10 0"
DIV_OUTPUT=$(echo "DIV 10 0" | ./calclang 2>&1 | grep -E "(Error|^>)" | head -n 2 | sed 's/^> //' | tr '\n' ' ')
echo "   $DIV_OUTPUT"
echo ""

# Demo multiple operations
echo "6. Multiple operations:"
echo "   ADD 100 50 = $(echo "ADD 100 50" | ./calclang | grep '^>' | head -n 1 | sed 's/^> //')"
echo "   SUB 100 50 = $(echo "SUB 100 50" | ./calclang | grep '^>' | head -n 1 | sed 's/^> //')"
echo "   MUL 12 8   = $(echo "MUL 12 8" | ./calclang | grep '^>' | head -n 1 | sed 's/^> //')"
echo "   DIV 96 8   = $(echo "DIV 96 8" | ./calclang | grep '^>' | head -n 1 | sed 's/^> //')"
echo ""

echo "Demo complete!"
echo ""
echo "To use CALCLANG interactively, run: ./calclang"
