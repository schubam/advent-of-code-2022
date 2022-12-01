package adventofcode2022

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(path string) string {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	return string(input)
}
