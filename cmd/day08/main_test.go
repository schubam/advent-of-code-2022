package main

import (
	"testing"
)

const input = `30373
25512
65332
33549
35390`

func Test_SolvePart2(t *testing.T) {
	want := 8
	got := SolveV2(input)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}

func Test_SolvePart1(t *testing.T) {
	want := 21
	got := SolveV1(input)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}
