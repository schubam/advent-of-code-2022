package main

import "testing"

func Test_parseData(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	want := []*Elf{
		{calories: []int{1000, 2000, 3000}},
		{calories: []int{4000}},
		{calories: []int{5000, 6000}},
		{calories: []int{7000, 8000, 9000}},
		{calories: []int{10000}},
	}

	got, err := parseData(input)
	if err != nil {
		t.Errorf("error parsing data: %s", err)
	}

	if len(got) != len(want) {
		t.Errorf("error: len(got)=%d, len(want)=%d", len(got), len(want))
	}

	for i, n := range got {
		if !n.isEqual(want[i]) {
			t.Errorf("want:%d, got:%d", want[i], n)
		}
	}
}
