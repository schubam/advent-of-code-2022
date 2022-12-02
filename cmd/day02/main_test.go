package main

import "testing"

func Test_SolvePart1(t *testing.T) {
	input := `A Y
B X
C Z`
	want := 15
	got := Solve(input, ScorePart1)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}

func Test_SolvePart2(t *testing.T) {
	input := `A Y
B X
C Z`
	want := 12
	got := Solve(input, ScorePart2)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}
