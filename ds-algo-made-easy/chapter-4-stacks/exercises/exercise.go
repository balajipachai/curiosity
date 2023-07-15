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

// The function `operatorPrecedence` returns the precedence level of a given operator.
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

// The function takes two integers and an operator as input and performs the corresponding operation on
// the integers.
func stackOperatorOperation(secondTopElement, topElement int, operator string) int {
	switch operator {
	case "*":
		return secondTopElement * topElement
	case "/":
		return secondTopElement / topElement
	case "-":
		return secondTopElement - topElement
	case "+":
		return secondTopElement + topElement
	default:
		return 0
	}
}

// The function `getIntFromAscii` takes an ASCII value as input and returns the corresponding integer
// value.
func getIntFromAscii(ascii int) int {
	switch ascii {
	case 48:
		return 0
	case 49:
		return 1
	case 50:
		return 2
	case 51:
		return 3
	case 52:
		return 4
	case 53:
		return 5
	case 54:
		return 6
	case 55:
		return 7
	case 56:
		return 8
	case 57:
		return 9
	default:
		return -1

	}
}

/*
* Algorithm:
* 1 Scan the Postfix string from left to right.
* 2 Initialize an empty stack.
* 3 Repeat steps 4 and 5 till all the characters are scanned.
* 4 If the scanned character is an operand, push it onto the stack.
* 5 If the scanned character is an operator, and if the operator is a unary operator, then
* pop an element from the stack. If the operator is a binary operator, then pop two elements from the stack. After popping the elements, apply the * operator to those popped elements. Let the result of this operation be retVal onto the stack.
* 6 After all characters are scanned, we will have only one element in the stack.
* 7 Return top of the stack as result.
 */
func PostfixEvaluation(dStack *stacks.DynamicStack, input string) int {
	for _, v := range input {
		if !isRuneASymbol(v) && !isRuneAnOperator(v) {
			element := getIntFromAscii(int(v))
			dStack.Push(element)
		} else if isRuneAnOperator(v) {
			topElement := dStack.Pop()
			secondTopElement := dStack.Pop()
			result := stackOperatorOperation(secondTopElement, topElement, string(v))
			dStack.Push(result)
		}
	}
	return dStack.TopElement()
}

func infixEvaluationHelper(operatorStack, operandStack *stacks.DynamicStack) {
	operator := operatorStack.TopElement()
	operatorStack.Pop()
	operand1 := operandStack.TopElement()
	operandStack.Pop()
	operand2 := operandStack.TopElement()
	operandStack.Pop()
	result := stackOperatorOperation(operand2, operand1, string(operator))
	operandStack.Push(result)
}

/*
Algorithm
If character exists to be read:

1. If character is operand push on the operand stack, if character is (, push on the operator stack.
2. Else if character is operator
 1. While the top of the operator stack is not of smaller precedence than this character.
 2. Pop operator from operator stack.
 3. Pop two operands (op1 and op2) from operand stack.
 4. Store op1 op op2 on the operand stack back to 2.1.

3. Else if character is ), do the same as 2.2 - 2.4 till you encounter (.

	Else (no more character left to read):

- Pop operators untill operator stack is not empty.
- Pop top 2 operands and push op1 op op2 on the operand stack.
return the top value from operand stack.

Source: https://stackoverflow.com/questions/13421424/how-to-evaluate-an-infix-expression-in-just-one-scan-using-stacks
*/
func InfixEvaluationUsingOnePass(operatorStack, operandStack *stacks.DynamicStack, input string) int {
	for _, v := range input {
		if !isRuneASymbol(v) && !isRuneAnOperator(v) {
			// If character is operand push on the operand stack
			operandStack.Push(getIntFromAscii(int(v)))
		} else if v == '(' { // if character is (, push on the operator stack
			operatorStack.Push(int(v))
		} else if v == ')' { // if character is ), do the same as 2.2 - 2.4 till you encounter (
			for operatorStack.TopElement() != '(' {
				infixEvaluationHelper(operatorStack, operandStack)
			}
		} else if isRuneAnOperator(v) {
			/*
				2. Else if character is operator
					 1. While the top of the operator stack is not of smaller precedence than this character.
					 2. Pop operator from operator stack.
					 3. Pop two operands (op1 and op2) from operand stack.
					 4. Store op1 op op2 on the operand stack back to 2.1.
			*/
			if !operatorStack.IsEmpty() {
				for operatorPrecedence(string(operatorStack.TopElement())) > operatorPrecedence(string(v)) {
					infixEvaluationHelper(operatorStack, operandStack)
				}
			}

		}
	}

	for !operatorStack.IsEmpty() {
		infixEvaluationHelper(operatorStack, operandStack)
	}
	return operandStack.TopElement()
}
