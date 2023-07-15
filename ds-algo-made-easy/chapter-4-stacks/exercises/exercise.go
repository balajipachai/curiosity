package exercises

import (
	"fmt"

	"example.com/stacks"
)

const (
	STACK_SIZE   = 10
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorCyan    = "\033[36m"
	colorYellow  = "\033[33m"
)

// The function `isRuneASymbol` checks if a given rune is one of the symbols
// `(`, `)`, `[`, `]`, `{`, or `}`.
func isRuneASymbol(c rune) bool {
	switch c {
	case 40:
		return true
	case 41:
		return true
	case 91:
		return true
	case 93:
		return true
	case 123:
		return true
	case 125:
		return true
	default:
		return false
	}
}

// The function checks if a given rune is an operator (+, -, *, or /) and returns true if it is,
// otherwise it returns false.
// 42 => *
// 43 => +
// 45 => -
// 47 => /
func isRuneAnOperator(c rune) bool {
	switch c {
	case 42:
		return true
	case 43:
		return true
	case 45:
		return true
	case 47:
		return true
	default:
		return false

	}
}

func operatorPrecedence(operator string) int {
	switch operator {
	case "^":
		return 3
	case "/", "*":
		return 2
	case "+", "-":
		return 1
	default:
		return -1

	}
}

// The IsSymbolBalanced function checks if a given string has balanced symbols (parentheses, brackets,
// and curly braces).
func IsSymbolBalanced(dStack *stacks.DynamicStack, input string) bool {
	/*
		ASCII Values for Symbols are as follows:
		'()' = 40, 41
		'[]' = 91, 93
		'{}' = 123, 125
	*/
	for _, v := range input {
		if isRuneASymbol(v) {
			fmt.Println(colorYellow + "\n\tPrint stack elements" + colorReset)
			dStack.Print(true)
			if v == 40 || v == 91 || v == 123 {
				fmt.Println("\tpush")
				fmt.Printf("\t%c", v)
				dStack.Push(int(v))
			} else {
				if dStack.IsEmpty() {
					fmt.Println(colorRed + "\tUnbalanced symbols" + colorReset)
					return false
				} else {
					fmt.Printf(colorCyan+"\tTopElement: %c", dStack.TopElement())
					poppedElement := dStack.Pop()
					fmt.Printf("\n\tpop\t%c", poppedElement)
					fmt.Println(colorReset)
					if int(v)-poppedElement > 2 {
						return false
					}
				}
			}
		}
	}
	if !dStack.IsEmpty() {
		fmt.Println(colorRed + "\tUnbalanced symbols" + colorReset)
		return false
	}
	fmt.Println(colorCyan + "\tBalanced symbols" + colorReset)
	return true
}

/*
The function takes an infix expression as input and converts it to postfix notation using a stack.
ASCII Value for Operators is as follows:

	42 => *
	43 => +
	45 => -
	47 => /
	A * B- (C + D) + E
*/
func InfixToPostfix(dStack *stacks.DynamicStack, input string) string {
	postfixExpression := ""
	for _, v := range input {
		// If 'v' is an operand push into the stack
		if !isRuneASymbol(v) && !isRuneAnOperator(v) {
			postfixExpression += string(v)
		} else if v == '(' { // if 'v' is left parentheses, push into the stack
			dStack.Push(int(v))
		} else if v == ')' { // if 'v' is right parentheses, pop & add to postfixExpression, until left parentheses is not encountered
			for dStack.TopElement() != '(' {
				postfixExpression += string(dStack.TopElement())
				dStack.Pop()
			}
			dStack.Pop()
		} else {
			// if 'v' is an operator, pop and add to postfixExpression until stack is not empty and operator precedence of v <= operator precedence of top stack element
			for !dStack.IsEmpty() && operatorPrecedence(string(v)) <= operatorPrecedence(string(dStack.TopElement())) {
				postfixExpression += string(dStack.TopElement())
				dStack.Pop()
			}
			dStack.Push(int(v))
		}
	}
	// Pop all remaining elements from the stack
	for !dStack.IsEmpty() {
		postfixExpression += string(dStack.TopElement())
		dStack.Pop()
	}
	return postfixExpression
}
