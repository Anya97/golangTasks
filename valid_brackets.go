package main

import "fmt"

func isBracketValid(s string) bool {
	var stack []rune
	openingBrackets := map[rune]bool{'(': true, '{': true, '[': true}
	closingBrackets := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, symbol := range s {
		if isOpeningBrackets := openingBrackets[symbol]; isOpeningBrackets {
			stack = append(stack, symbol)
		} else if _, isClosingBrackets := closingBrackets[symbol]; isClosingBrackets {
			if len(stack) == 0 {
				return false
			}
			openingBracket := closingBrackets[symbol]
			if stack[len(stack)-1] == openingBracket {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Printf("%t\n", isBracketValid(")("))
	fmt.Printf("%t\n", isBracketValid("(){}[]"))
	fmt.Printf("%t\n", isBracketValid("([)]"))
	fmt.Printf("%t\n", isBracketValid("({[()]})[]"))
	fmt.Printf("%t\n", isBracketValid("("))
	fmt.Printf("%t\n", isBracketValid("(())()"))
	fmt.Printf("%t\n", isBracketValid("())("))
	fmt.Printf("%t\n", isBracketValid("()))))((((()"))
	fmt.Printf("%t\n", isBracketValid("1+(3*6)"))
	fmt.Printf("%t\n", isBracketValid(")1+(3*6)"))
}
