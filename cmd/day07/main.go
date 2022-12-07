package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/schubam/advent-of-code-2022/tree"
)

func parseData(input io.Reader) *tree.Tree {
	scanner := bufio.NewScanner(input)

	fs := tree.NewTree("/")

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println("line:", line)

		components := strings.Split(line, " ")

		if strings.HasPrefix(line, "$") {
			cmd := components[1]
			switch cmd {
			case "cd":
				if len(components) > 2 {
					arg := components[2]
					fs.Traverse(arg)
				}
			case "ls":
				continue
			}
		} else {
			dirOrSize := components[0]
			name := components[1]
			if strings.HasPrefix(dirOrSize, "dir") {
				fs.CreateDir(name)
			} else {
				size, _ := strconv.Atoi(dirOrSize)
				fs.CreateFile(name, int32(size))
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	//fmt.Println("parsing done.")

	return fs
}

func SolveV1(input io.Reader) int32 {
	t := parseData(input)
	//t.Print()

	return tree.DirectorySizes(t.Root)
}

func SolveV2(input io.Reader) int32 {
	t := parseData(input)
	t.Print()

	return t.FindSmallestDir()
}

func main() {
	input, _ := os.Open("input.txt")
	fmt.Println(SolveV1(input))
	input, _ = os.Open("input.txt")
	fmt.Println(SolveV2(input))
}
