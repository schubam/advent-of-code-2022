package main

import "testing"

func Test_SolvePart2(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	want := 70
	got := SolveV2(input)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}

func Test_SolvePart1(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	want := 157
	got := SolveV1(input)
	if got != want {
		t.Errorf("error: got=%d,want=%d\n", got, want)
	}
}

func Test_FindDups(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp"
	first := removeDuplicates(input[0 : len(input)/2])
	second := removeDuplicates(input[len(input)/2:])
	words := []string{first, second}
	want := []rune{'p'}
	got := FindDups(words)

	if len(got) != len(want) {
		t.Errorf("error: got=%c,want=%c\n", got, want)
	} else {
		for i, v := range want {
			if v != got[i] {
				t.Errorf("error: got=%c,want=%c\n", got[i], v)
			}
		}
	}
}
