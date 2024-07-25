// the program is sorting the numbers provided in the terminal
package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	SliceOfStrings := ParseTerminal()
	SliceOfNumbers := CreateSlice(SliceOfStrings)
	SortedSliceOfNumbers := SortSlice(SliceOfNumbers)
	fmt.Println("The sorted results are: ", SortedSliceOfNumbers)
}

func ParseTerminal() []string { 
	flag.Parse()
	sliceOfStrings := flag.Args()
	return sliceOfStrings 
}

func CreateSlice(SliceOfStrings []string) []int {
	var SliceOfNumbers []int
	for i := range SliceOfStrings {
		value, err := strconv.Atoi(SliceOfStrings[i])
		if err != nil {
			panic(err)
		}
		SliceOfNumbers = append(SliceOfNumbers, value)
	}
	fmt.Println("The input values are: ", SliceOfNumbers)
	return SliceOfNumbers
}

func SortSlice(SliceOfNumbers []int) []int {
	var sortedSliceOfNumbers []int
	for j := range SliceOfNumbers {
		if j == len(SliceOfNumbers)-1 {
			break
		} else if SliceOfNumbers[j] > SliceOfNumbers[j+1] {
			SliceOfNumbers[j], SliceOfNumbers[j+1] = SliceOfNumbers[j+1], SliceOfNumbers[j]
		}
		sortedSliceOfNumbers = SliceOfNumbers
		for j := range SliceOfNumbers {
			if j == len(SliceOfNumbers)-1 {
				break
			} else if SliceOfNumbers[j] > SliceOfNumbers[j+1] {
				SliceOfNumbers[j], SliceOfNumbers[j+1] = SliceOfNumbers[j+1], SliceOfNumbers[j]
			}
			sortedSliceOfNumbers = SliceOfNumbers
		}
	}
	return sortedSliceOfNumbers
}
