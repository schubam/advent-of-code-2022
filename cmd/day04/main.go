package main

import (
	"fmt"
	"strconv"
	"strings"

	adventofcode2022 "github.com/schubam/advent-of-code-2022"
)

type Pair struct {
	one CleaningJob
	two CleaningJob
}

// a: 2-3
// b: 4-5
func NewPair(a, b string) Pair {
	return Pair{one: parseCleaningJob(a), two: parseCleaningJob(b)}
}

func (p Pair) isAnyContainingTheOther() bool {
	return p.one.isFullyContains(p.two) || p.two.isFullyContains(p.one)
}

func (p Pair) isAnyOverlapping() bool {
	return p.one.isOverlapping(p.two)
}

type CleaningJob struct {
	low  int
	high int
}

func (p CleaningJob) isFullyContains(other CleaningJob) bool {
	lower := p.low <= other.low
	higher := p.high >= other.high

	return lower && higher
}

func (p CleaningJob) isOverlapping(other CleaningJob) bool {
	r := createNumbers(p.low, p.high)
	o := createNumbers(other.low, other.high)
	for _, v := range r {
		for _, w := range o {
			if v == w {
				return true
			}
		}
	}

	return false
}

func createNumbers(lo int, hi int) []int {
	s := make([]int, hi-lo+1)

	for i := range s {
		s[i] = i + lo
	}

	return s
}

func parseCleaningJob(a string) CleaningJob {
	var res CleaningJob

	numbers := strings.Split(a, "-")

	v, err := strconv.Atoi(numbers[0])
	if err != nil {
		fmt.Printf("can't convert to number:%s\n", err)
	}
	res.low = v

	v, err = strconv.Atoi(numbers[1])
	if err != nil {
		fmt.Printf("can't convert to number:%s\n", err)
	}
	res.high = v

	return res
}

func SolveV1(input string) int {
	var score int
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) > 0 {
			comps := strings.Split(line, ",")
			pair := NewPair(comps[0], comps[1])
			if pair.isAnyContainingTheOther() {
				score += 1
			}
		}
	}

	return score
}

func SolveV2(input string) int {
	var score int
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) > 0 {
			comps := strings.Split(line, ",")
			pair := NewPair(comps[0], comps[1])
			if pair.isAnyOverlapping() {
				score += 1
			}
		}
	}

	return score
}

func main() {
	input := adventofcode2022.ReadFile("input.txt")
	fmt.Println(SolveV1(input))
	fmt.Println(SolveV2(input))
}
