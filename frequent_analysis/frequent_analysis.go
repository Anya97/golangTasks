package main

import (
	"fmt"
	"sort"
	"strings"
)

type keyValue struct {
	Key   string
	Value int
}

func main() {
	fmt.Println(Top10(`cat and dog, one dog,two cats and one man  `))
}

func Top10(word string) []string {
	sliceOfWords := strings.Split(word, " ")
	mapCounter := make(map[string]int)

	for _, j := range sliceOfWords {
		if len(j) == 0 {
			continue
		}
		mapCounter[j] += 1
	}

	var result []keyValue

	for key, value := range mapCounter {
		result = append(result, keyValue{key, value})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Value == result[j].Value {
			return result[i].Key > result[j].Key
		}
		return result[i].Value > result[j].Value
	})

	var resultSlice []string
	for i, keyValue := range result {
		if i == 10 {
			break
		}
		resultSlice = append(resultSlice, keyValue.Key)
	}

	return resultSlice
}
