package main

import (
	"strings"
	"testing"
)

const input = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func Test_SolvePart2(t *testing.T) {
	want := int32(24933642)
	reader := strings.NewReader(input)
	got := SolveV2(reader)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}

func Test_SolvePart1(t *testing.T) {
	want := int32(95437)
	reader := strings.NewReader(input)
	got := SolveV1(reader)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}
