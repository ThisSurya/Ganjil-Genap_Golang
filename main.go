package main

import (
	"fmt"
	// "strings"
)

func main() {
	Number := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// rawResult := []string{}

	for i := range Number {
		if Number[i]%2 == 0 {
			fmt.Println(i, "is Even")
		} else {
			fmt.Println(i, "is Odd")
		}
	}

	// result := strings.Join(rawResult, ", ")
}
