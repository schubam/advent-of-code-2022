package main

import (
	"errors"
	"fmt"
	"github.com/schubam/advent-of-code-2022"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	index    int
	calories []int
}

func (e *Elf) isEqual(other *Elf) bool {
	if len(e.calories) != len(other.calories) {
		return false
	}

	for i, n := range e.calories {
		if n != other.calories[i] {
			return false
		}
	}

	return true
}

func (e *Elf) totalCalories() int {
	var sum int
	for _, c := range e.calories {
		sum += c
	}

	return sum
}

func parseData(input string) ([]*Elf, error) {
	elves := []*Elf{}

	lines := strings.Split(input, "\n")

	if len(lines) < 1 {
		return []*Elf{}, errors.New("no input lines found")
	}

	idx := 0
	elf := &Elf{index: idx}

	for _, line := range lines {
		l := strings.TrimSpace(line)
		if l == "" {
			elves = append(elves, elf)
			idx += 1
			elf = &Elf{index: idx}
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			fmt.Printf("error converting line to number: %s\n", l)
		}
		elf.calories = append(elf.calories, n)
	}

	elves = append(elves, elf)
	return elves, nil
}

func main() {
	input := adventofcode2022.ReadFile("input.txt")

	elves, err := parseData(input)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].totalCalories() > elves[j].totalCalories()
	})

	fmt.Println("part 1: ", elves[0].totalCalories())
	fmt.Println("part 2: ", elves[0].totalCalories()+elves[1].totalCalories()+elves[2].totalCalories())
}
