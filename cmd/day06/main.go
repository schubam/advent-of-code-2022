package main

import (
	"fmt"

	adventofcode2022 "github.com/schubam/advent-of-code-2022"
)

func hasDuplicates(word string) bool {
	occ := map[rune]bool{}
	for _, c := range word {
		if _, ok := occ[c]; !ok {
			occ[c] = true
		} else {
			return true
		}
	}
	return false
}

func SolveV1(input string, window int) int {
	var result int

	for i := range input {
		arr := input[i : i+window]
		if !hasDuplicates(arr) {
			return i + window
		}
	}

	return result
}

func main() {
	input := adventofcode2022.ReadFile("input.txt")
	fmt.Println(SolveV1(input, 4))
	fmt.Println(SolveV1(input, 14))
}
