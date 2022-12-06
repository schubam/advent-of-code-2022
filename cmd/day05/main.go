package main

import (
	"fmt"
	"strconv"
	"strings"

	adventofcode2022 "github.com/schubam/advent-of-code-2022"
)

type Stack []string

func (s Stack) Push(str string) Stack {
	return append(s, str)
}

func (s Stack) Pop() (Stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s Stack) Print() {
	for _, ele := range s {
		fmt.Printf("%s", ele)
	}
	fmt.Println()
}

type State struct {
	columns []*Stack
}

func (s *State) Print() {
	for i, stack := range s.columns {
		fmt.Printf("[%d]:", i+1)
		stack.Print()
	}
}

func (s *State) execute(ins Instruction) {
	var counter int
	for {
		counter++
		from := ins[1] - 1
		to := ins[2] - 1
		fromPtr := s.columns[from]
		toPtr := s.columns[to]

		popped, ele := fromPtr.Pop()
		s.columns[from] = &popped

		pushed := toPtr.Push(ele)
		s.columns[to] = &pushed

		if counter == ins[0] {
			counter = 0
			break
		}
	}
}

func (s *State) execute2(ins Instruction) {
	var counter int
	from := ins[1] - 1
	to := ins[2] - 1

	var eles []string
	for {
		counter++
		popped, ele := s.columns[from].Pop()
		s.columns[from] = &popped

		eles = append(eles, ele)

		if counter == ins[0] {
			counter = 0
			break
		}
	}
	eles = reverseArray(eles)

	for _, e := range eles {
		pts := s.columns[to].Push(e)
		s.columns[to] = &pts
	}
}

func (s State) topElements() string {
	var res []string
	for _, ptr := range s.columns {
		stack := *ptr
		res = append(res, stack[len(stack)-1])
	}
	return strings.Join(res, "")
}

func parseHeader(lines []string) State {
	rowSize := (len(lines[0]) + 1) / 4
	//fmt.Printf("num columns=%d\n", rowSize)

	stacks := []*Stack{}
	for i := 0; i < rowSize; i++ {
		stacks = append(stacks, &Stack{})
	}
	res := State{columns: stacks}

	for _, line := range lines {
		if line == "" {
			break
		}
		//fmt.Println(line)
		for i := 0; i < rowSize; i++ {
			crate := string(line[4*i+1])
			if crate != " " {
				st := (*res.columns[i]).Push(crate)
				res.columns[i] = &st
			}
		}
	}

	// remove last element and then reverse the array
	for i, v := range res.columns {
		stack := *v
		st := reverseArray(stack[:len(stack)-1])
		res.columns[i] = &st
	}

	return res
}

func reverseArray(arr Stack) Stack {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

type Instruction [3]int

func parseInstructions(lines []string) []Instruction {
	res := []Instruction{}
	for _, line := range lines {
		if strings.HasPrefix(line, "move") {
			cs := strings.Split(line, " ")
			//fmt.Printf("components: %+v\n", cs)

			x, _ := strconv.Atoi(cs[1])
			y, _ := strconv.Atoi(cs[3])
			z, _ := strconv.Atoi(cs[5])

			res = append(res, Instruction{x, y, z})
		}
	}
	return res
}

func executeInstructions(state *State, instrs []Instruction) string {
	for _, ins := range instrs {
		state.execute(ins)
	}

	return state.topElements()
}

func executeInstructions2(state *State, instrs []Instruction) string {
	for _, ins := range instrs {
		state.execute2(ins)
	}

	return state.topElements()
}

func SolveV1(input string) string {
	lines := strings.Split(input, "\n")

	state := parseHeader(lines)

	instructions := parseInstructions(lines)
	//fmt.Println(instructions)

	result := executeInstructions(&state, instructions)
	state.Print()

	return result
}

func SolveV2(input string) string {
	lines := strings.Split(input, "\n")

	state := parseHeader(lines)

	instructions := parseInstructions(lines)
	//fmt.Println(instructions)

	result := executeInstructions2(&state, instructions)
	state.Print()

	return result
}

func main() {
	input := adventofcode2022.ReadFile("input.txt")
	fmt.Println(SolveV1(input))
	fmt.Println(SolveV2(input))
}
