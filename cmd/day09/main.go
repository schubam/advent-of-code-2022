package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	adventofcode2022 "github.com/schubam/advent-of-code-2022"
)

type Vec struct {
	x int
	y int
}

func SolveV1(input string) int {
	chain := []*Vec{{0, 0}, {0, 0}}
	head := chain[0]
	tail := chain[len(chain)-1]
	visits := map[Vec]int{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			break
		}
		components := strings.Split(line, " ")
		direction := components[0]
		amount, _ := strconv.Atoi(components[1])

		dir := &Vec{0, 0}
		switch direction {
		case "U":
			dir.y = 1
		case "D":
			dir.y = -1
		case "L":
			dir.x = -1
		case "R":
			dir.x = 1
		}

		for i := 0; i < amount; i++ {
			head.x += dir.x
			head.y += dir.y

			for j := 0; j < len(chain)-1; j++ {
				cur := chain[j]
				next := chain[j+1]

				fmt.Println(*cur, *next)

				if !adjacent(cur, next) {
					if cur.x > next.x {
						next.x += 1
					}
					if cur.x < next.x {
						next.x -= 1
					}
					if cur.y > next.y {
						next.y += 1
					}
					if cur.y < next.y {
						next.y -= 1
					}
				}
				fmt.Println(*cur, *next)
				fmt.Println()

				visits[*tail] += 1
			}
		}
	}

	return len(visits)
}

func adjacent(a, b *Vec) bool {
	return math.Abs(float64(a.x)-float64(b.x)) <= 1 && math.Abs(float64(a.y)-float64(b.y)) <= 1
}

func SolveV2(input string) int {
	return 0
}

func main() {
	input := adventofcode2022.ReadFile("input.txt")
	fmt.Println(SolveV1(input))
	fmt.Println(SolveV2(input))
}
