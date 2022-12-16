package main

import (
	"testing"
)

func Test_SolvePart2(t *testing.T) {
	input := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	want := 36
	got := SolveV2(input, false)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}

func xTest_SolvePart1(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	want := 13
	got := SolveV1(input, false)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}
