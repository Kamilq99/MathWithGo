package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// Function that calculates the value of a mathematical expression with parentheses.
func evaluate(expression string) (int, error) {
	// Check if parentheses are balanced
	if !isValidParentheses(expression) {
		return 0, errors.New("mismatched opening and closing parentheses")
	}

	// Use a stack to evaluate the expression
	return eval(expression)
}

// Function to check if the number of opening and closing parentheses is balanced.
func isValidParentheses(expression string) bool {
	balance := 0
	for _, char := range expression {
		if char == '(' {
			balance++
		} else if char == ')' {
			balance--
		}
		// If a closing parenthesis appears before an opening one at any point
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}

// Function to handle calculations using a stack.
func eval(expression string) (int, error) {
	var nums []int
	var ops []rune

	// Function to define operator precedence
	priority := func(op rune) int {
		if op == '+' || op == '-' {
			return 1
		}
		if op == '*' || op == '/' {
			return 2
		}
		return 0
	}

	// Function to apply the top operator to the top two numbers in the stack
	applyOp := func() {
		if len(nums) < 2 || len(ops) == 0 {
			return
		}
		b := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		a := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		var result int
		switch op {
		case '+':
			result = a + b
		case '-':
			result = a - b
		case '*':
			result = a * b
		case '/':
			result = a / b
		}
		nums = append(nums, result)
	}

	// Main loop to process the expression
	for i := 0; i < len(expression); i++ {
		char := rune(expression[i])

		// If the character is a digit, parse the full number
		if unicode.IsDigit(char) {
			start := i
			for i < len(expression) && unicode.IsDigit(rune(expression[i])) {
				i++
			}
			num, _ := strconv.Atoi(expression[start:i])
			nums = append(nums, num)
			i--
		} else if char == '(' {
			ops = append(ops, char)
		} else if char == ')' {
			// Apply all operators until we encounter the opening parenthesis
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				applyOp()
			}
			ops = ops[:len(ops)-1] // Remove the opening parenthesis
		} else if char == '+' || char == '-' || char == '*' || char == '/' {
			// Apply operators with higher or equal precedence
			for len(ops) > 0 && priority(ops[len(ops)-1]) >= priority(char) {
				applyOp()
			}
			ops = append(ops, char)
		}
	}

	// Apply any remaining operators
	for len(ops) > 0 {
		applyOp()
	}

	return nums[0], nil
}

func main() {
	expression := "3+(2*2)-(1+2)"
	result, err := evaluate(expression)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
