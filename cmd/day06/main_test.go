package main

import "testing"

func Test_SolvePart2(t *testing.T) {
	input := `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
	want := 19
	got := SolveV1(input, 14)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}

func Test_SolvePart1(t *testing.T) {
	input := `bvwbjplbgvbhsrlpgdmjqwftvncz`
	want := 5
	got := SolveV1(input, 4)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}
