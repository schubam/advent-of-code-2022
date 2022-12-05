package main

import "testing"

func Test_SolvePart2(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	want := 4
	got := SolveV2(input)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}

func Test_SolvePart1(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	want := 2
	got := SolveV1(input)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}

func Test_PairIsFullyContains(t *testing.T) {
	input := CleaningJob{low: 2, high: 8}
	other := CleaningJob{low: 3, high: 7}
	want := true
	got := input.isFullyContains(other)

	if got != want {
		t.Errorf("input:%+v, other:%+v,got:%t,want:%t", input, other, got, want)
	}
}
