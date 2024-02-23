package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Unpack("a4bc2d5e"))
	fmt.Println(Unpack("abcd"))
	fmt.Println(Unpack("3abc"))
	fmt.Println(Unpack("45"))
	fmt.Println(Unpack("aaa10b"))
	fmt.Println(Unpack("aaa0b"))
	fmt.Println(Unpack(""))
	fmt.Println(Unpack("d\n5abc"))
}

func Unpack(str string) (string, error) {
	listRunes := []rune(str)
	if len(str) == 0 {
		return "", nil
	}

	var resultStr strings.Builder

	for i := 0; i < len(listRunes)-1; i++ {
		if (i == 0 && unicode.IsDigit(listRunes[i])) || (unicode.IsDigit(listRunes[i]) && unicode.IsDigit(listRunes[i+1])) {
			return "", errors.New("некорректная строка")
		}

		if !unicode.IsDigit(listRunes[i]) && !unicode.IsDigit(listRunes[i+1]) {
			resultStr.WriteString(string(listRunes[i]))
		} else if !unicode.IsDigit(listRunes[i]) {
			num, _ := strconv.Atoi(string(listRunes[i+1]))
			resultStr.WriteString(strings.Repeat(string(listRunes[i]), num))
		}
	}
	if !unicode.IsDigit(listRunes[len(listRunes)-1]) {
		resultStr.WriteString(string(listRunes[len(listRunes)-1]))
	}

	return resultStr.String(), nil
}
