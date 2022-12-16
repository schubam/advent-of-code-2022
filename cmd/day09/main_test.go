package main

import (
	"testing"
)

const input = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func xTest_SolvePart2(t *testing.T) {
	want := -1
	got := SolveV2(input)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}

func Test_SolvePart1(t *testing.T) {
	want := 13
	got := SolveV1(input)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}
