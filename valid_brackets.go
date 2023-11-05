package main

import "fmt"

func isBracketValid(s string) bool {
	bracketsCount := 0
	for _, v := range s {
		if string(v) == "(" {
			bracketsCount += 1
		} else if string(v) == ")" {
			bracketsCount -= 1
		}
		if bracketsCount < 0 {
			return false
		}
	}
	return bracketsCount == 0
}

func main() {
	fmt.Printf("%t\n", isBracketValid(")(")) // false
	fmt.Printf("%t\n", isBracketValid("("))  // false
	fmt.Printf("%t\n", isBracketValid("(())()"))
	fmt.Printf("%t\n", isBracketValid("())("))
	fmt.Printf("%t\n", isBracketValid("()))))((((()"))
	fmt.Printf("%t\n", isBracketValid("1+(3*6)"))
	fmt.Printf("%t\n", isBracketValid(")1+(3*6)"))
}
