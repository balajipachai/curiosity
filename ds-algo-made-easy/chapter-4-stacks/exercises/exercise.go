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
