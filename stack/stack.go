package stack

import "fmt"

type Stack []string

func (s Stack) Push(str string) Stack {
	return append(s, str)
}

func (s Stack) Pop() (Stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s Stack) Print() {
	for _, ele := range s {
		fmt.Printf("%s", ele)
	}
	fmt.Println()
}
