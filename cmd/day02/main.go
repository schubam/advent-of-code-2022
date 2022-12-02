package main

import (
	"fmt"
	"strings"

	adventofcode2022 "github.com/schubam/advent-of-code-2022"
)

const (
	ROCK     = "rock"
	PAPER    = "paper"
	SCISSORS = "scissors"
)

var POINTS = map[string]int{
	ROCK: 1, PAPER: 2, SCISSORS: 3,
}

var MAPPING = map[string]string{
	"A": ROCK,
	"X": ROCK,
	"B": PAPER,
	"Y": PAPER,
	"C": SCISSORS,
	"Z": SCISSORS,
}

func winOrLose(a string, b string) int {
	if a == b {
		return 3
	} else if a == ROCK && b == PAPER {
		return 6
	} else if a == PAPER && b == SCISSORS {
		return 6
	} else if a == SCISSORS && b == ROCK {
		return 6
	}
	return 0
}

func ScorePart1(first string, second string) int {
	a := MAPPING[first]
	b := MAPPING[second]

	return POINTS[b] + winOrLose(a, b)
}

func ScorePart2(first string, second string) int {
	a := MAPPING[first]
	var b string

	switch second {
	case "X":
		// lose
		if a == ROCK {
			b = SCISSORS
		} else if a == PAPER {
			b = ROCK
		} else if a == SCISSORS {
			b = PAPER
		}
	case "Y":
		// draw
		b = a
	case "Z":
		// win
		if a == ROCK {
			b = PAPER
		} else if a == PAPER {
			b = SCISSORS
		} else if a == SCISSORS {
			b = ROCK
		}
	}
	return POINTS[b] + winOrLose(a, b)
}

func Solve(input string, scorer func(a, b string) int) int {
	lines := strings.Split(input, "\n")
	var points int

	for i, l := range lines {
		selections := strings.Split(l, " ")
		if len(selections) == 2 {
			points += scorer(selections[0], selections[1])
		} else {
			fmt.Printf("error on line %d: `%s`\n", i, l)
		}
	}
	return points
}

func main() {
	input := adventofcode2022.ReadFile("input.txt")
	fmt.Println(Solve(input, ScorePart1))
	fmt.Println(Solve(input, ScorePart2))
}
