package main

import (
	"fmt"
	"strconv"
	"strings"

	adventofcode2022 "github.com/schubam/advent-of-code-2022"
)

type Forest struct {
	grid   []int
	width  int
	height int
}

func (f *Forest) get(x, y int) int {
	idx := f.width*y + x
	//fmt.Printf("get: len=%d,lastIdx=%d,x=%d,y=%d => idx=%d\n", len(f.grid), len(f.grid)-1, x, y, idx)
	return f.grid[idx]
}

func (f *Forest) castRay(x, y int, direction [2]int) []int {
	res := []int{}

	for {
		x += direction[0]
		y += direction[1]
		if x >= f.width || y >= f.height || x < 0 || y < 0 {
			break
		}
		res = append(res, f.get(x, y))
	}

	return res
}

func (f *Forest) isVisible(i, j int) bool {
	t := f.get(i, j)
	//fmt.Printf("visible: i=%d,j=%d,t=%d\n", i, j, t)

	directions := [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	for _, d := range directions {
		nodes := f.castRay(i, j, d)
		if checkNodes(t, nodes) {
			return true
		}
	}

	return false
}

func (f *Forest) scenicAmount(i, j int) int {
	res := 1
	t := f.get(i, j)
	//fmt.Printf("visible: i=%d,j=%d,t=%d\n", i, j, t)

	viewingDistances := []int{}

	directions := [][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	for _, d := range directions {
		nodes := f.castRay(i, j, d)
		var dist int

		for _, n := range nodes {
			if n < t {
				dist++
			} else if n >= t {
				dist++
				break
			}
		}
		viewingDistances = append(viewingDistances, dist)
	}

	for _, n := range viewingDistances {
		res = res * n
	}
	fmt.Printf("i=%d,j=%d,t=%d,viewingDistances:%v,res=%d\n", i, j, t, viewingDistances, res)
	return res
}

func (f *Forest) numVisibleTrees() int {
	var sum int
	for i := 0; i < f.width; i++ {
		for j := 0; j < f.height; j++ {
			if f.isVisible(i, j) {
				//fmt.Printf("visible: i=%d,j=%d,t=%d\n", i, j, f.get(i, j))
				sum++
			}

		}
	}
	return sum
}

func (f *Forest) mostScenicTree() int {
	var max int
	for i := 0; i < f.width; i++ {
		for j := 0; j < f.height; j++ {
			s := f.scenicAmount(i, j)
			if s > max {
				max = s
			}
		}
	}
	return max
}

func checkNodes(size int, nodes []int) bool {
	var vis int
	for _, n := range nodes {
		if n < size {
			vis++
		}
	}
	return vis == len(nodes)
}

func NewForest(input string) *Forest {
	lines := strings.Split(input, "\n")

	var g []int
	for _, l := range lines {
		for _, c := range l {
			n, _ := strconv.Atoi(string(c))
			g = append(g, n)
		}
	}
	w := len(lines[0])
	f := &Forest{
		width:  w,
		height: len(g) / w,
		grid:   g,
	}

	//fmt.Println(f.width, f.height)
	return f
}

func SolveV1(input string) int {
	f := NewForest(input)

	return f.numVisibleTrees()
}

func SolveV2(input string) int {
	f := NewForest(input)

	return f.mostScenicTree()
}

func main() {
	input := adventofcode2022.ReadFile("input.txt")
	fmt.Println(SolveV1(input))
	fmt.Println(SolveV2(input))
}
