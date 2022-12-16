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

type found struct {
	vec *Vec
	idx int
}

func SolveV1(input string, shouldRender bool) int {
	return solve(input, 2, shouldRender)
}

func SolveV2(input string, shouldRender bool) int {
	return solve(input, 10, shouldRender)
}

func isAdjacent(a, b *Vec) bool {
	return math.Abs(float64(a.x)-float64(b.x)) <= 1 && math.Abs(float64(a.y)-float64(b.y)) <= 1
}

func render(chain []*Vec, visits map[Vec]int) {
	width := 80
	height := 50

	l := ""
	for i := -50; i < height-50; i++ {
		for j := -30; j < width-30; j++ {
			token := ""
			f := found{idx: -1}
			for idx, n := range chain {
				if n.x == i && n.y == j {
					f.vec = n
					f.idx = idx
					break
				}
			}

			switch {
			case f.idx == 0:
				token = "H"
			case f.idx == len(chain)-1:
				token = "T"
			case 1 <= f.idx && f.idx < len(chain)-1:
				token = fmt.Sprintf("%d", f.idx)
			default:
				token += "."
			}

			for v := range visits {
				if v.x == i && v.y == j {
					token = "#"
				}
			}
			l += token
		}
		l += fmt.Sprintf("\n")
	}
	fmt.Print("\033[2J", "\033[H", l)
}

func solve(input string, length int, shouldRender bool) int {
	chain := make([]*Vec, length)
	for i := 0; i < length; i++ {
		chain[i] = &Vec{0, 0}
	}
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
			if shouldRender {
				render(chain, visits)
			}
			head.x += dir.x
			head.y += dir.y

			for j := 0; j < len(chain)-1; j++ {
				cur := chain[j]
				next := chain[j+1]

				if !isAdjacent(cur, next) {
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
			}
			visits[*tail] += 1
		}
	}
	if shouldRender {
		render(chain, visits)
	}

	return len(visits)
}

func main() {
	input := adventofcode2022.ReadFile("input.txt")
	fmt.Println(SolveV1(input, false))
	fmt.Println(SolveV2(input, false))
}
