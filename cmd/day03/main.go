package main

import (
	"fmt"
	"strings"

	adventofcode2022 "github.com/schubam/advent-of-code-2022"
)

func removeDuplicates(s string) string {
	var result string
	occurrences := map[rune]bool{}

	for _, c := range s {
		if _, exists := occurrences[c]; !exists {
			occurrences[c] = true
			result = fmt.Sprintf("%s%c", result, c)
		}
	}

	return result
}

func FindDups(words []string) []rune {
	results := []rune{}
	occurrences := map[rune]int{}

	for _, word := range words {
		word = removeDuplicates(word)
		for _, r := range word {
			occurrences[r] = occurrences[r] + 1
		}
	}

	for k, v := range occurrences {
		if v == len(words) {
			results = append(results, k)
		}
	}

	return results
}

func CalculateScore(letters []rune) int {
	var sum int
	for _, l := range letters {
		switch {
		case l > 64 && l < 97:
			delta := int(l) - 38
			//fmt.Printf("adding %d for letter %c\n", delta, l)
			sum += delta
		case l > 96:
			delta := int(l) - 96
			//fmt.Printf("adding %d for letter %c\n", delta, l)
			sum += delta
		}
	}
	return sum
}

func SolveV1(input string) int {
	lines := strings.Split(input, "\n")
	var score int

	for _, l := range lines {

		first := l[0 : len(l)/2]
		second := l[len(l)/2:]
		words := []string{first, second}
		score += CalculateScore(FindDups(words))
	}

	return score
}

func SolveV2(input string) int {
	lines := strings.Split(input, "\n")
	var score int

	group := []string{}
	for i, l := range lines {
		group = append(group, l)
		if (i+1)%3 == 0 {
			dups := FindDups(group)
			score += CalculateScore(dups)
			group = []string{}
		}
	}

	return score
}

func main() {
	input := adventofcode2022.ReadFile("input.txt")
	fmt.Println(SolveV1(input))
	fmt.Println(SolveV2(input))
}
