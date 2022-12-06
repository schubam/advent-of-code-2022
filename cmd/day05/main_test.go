package main

import "testing"

func Test_SolvePart2(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	want := "MCD"
	got := SolveV2(input)
	if got != want {
		t.Errorf("error: got=%s,want=%s\n", got, want)
	}
}

func Test_SolvePart1(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	want := "CMZ"
	got := SolveV1(input)
	if got != want {
		t.Errorf("error: got=%s,want=%s\n", got, want)
	}
}

func Test_Stack(t *testing.T) {
	stack := Stack{"A", "B", "C"}
	want := "Z"
	wanted := Stack{"A", "B", "C"}

	stack = stack.Push("Z")
	stack, got := stack.Pop()

	if got != want {
		t.Errorf("error: got=%s,want=%s\n", got, want)
	}

	for i, e := range stack {
		if e != wanted[i] {
			t.Errorf("error: got=%s,want=%s,stack=%v,wanted=%v\n", got,
				wanted[i], stack, wanted)
		}
	}
}
