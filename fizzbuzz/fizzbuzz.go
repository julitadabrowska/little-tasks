package main

import (
	"fmt"
	"strconv"
)

func main() {
	outputSlice := SolveFizzBuzz(100)
	fmt.Println(outputSlice)
}

func SolveFizzBuzz(n int) []string {
	var outputSlice []string
	for i := 1; i <= n; i++ {
		if (i%3) == 0 && (i%5) == 0 {
			outputSlice = append(outputSlice, "FizzBuzz")
		} else if (i % 5) == 0 {
			outputSlice = append(outputSlice, "Buzz")
		} else if (i % 3) == 0 {
			outputSlice = append(outputSlice, "Fizz")
		} else {
			j := strconv.Itoa(i)

			outputSlice = append(outputSlice, j)
		}
	}
	return outputSlice //the return needs to be outside of the loop because we want the loop go all the iterations before we return the output
}
