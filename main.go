package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// AST represents an abstract syntax tree node
type AST interface{}

// Val represents a numeric value
type Val struct {
	value int
}

// Add represents an addition operation with two values
type Add struct {
	left  Val
	right Val
}

// Sub represents a subtraction operation with two values
type Sub struct {
	left  Val
	right Val
}

// Mul represents a multiplication operation with two values
type Mul struct {
	left  Val
	right Val
}

// Div represents a division operation with two values
type Div struct {
	left  Val
	right Val
}

// Mod represents a modulo operation with two values
type Mod struct {
	left  Val
	right Val
}

// Eval evaluates an AST node and returns its result
func Eval(ast AST) int {
	switch node := ast.(type) {
	case Add:
		return node.left.value + node.right.value
	case Sub:
		return node.left.value - node.right.value
	case Mul:
		return node.left.value * node.right.value
	case Div:
		return node.left.value / node.right.value
	case Mod:
		return node.left.value % node.right.value
	case Val:
		return node.value
	default:
		return 0
	}
}

// Parse converts a string expression into an operation struct
func Parse(input string) AST {
	parts := strings.Split(input, " ")
	left, _ := strconv.Atoi(parts[1])
	right, _ := strconv.Atoi(parts[2])

	switch parts[0] {
	case "ADD":
		return Add{Val{left}, Val{right}}
	case "SUB":
		return Sub{Val{left}, Val{right}}
	case "MUL":
		return Mul{Val{left}, Val{right}}
	case "DIV":
		return Div{Val{left}, Val{right}}
	case "MOD":
		return Mod{Val{left}, Val{right}}
	default:
		return nil
	}
}

func main() {
	fmt.Printf("CALCLANG\n\n")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if input == "" {
			continue
		}

		expr := Parse(input)
		if expr == nil {
			fmt.Println("Invalid expression")
			continue
		}
		fmt.Println(Eval(expr))
	}
}
