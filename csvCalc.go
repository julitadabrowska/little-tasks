// the program is openig a csv file and making a calculation written in the file
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	fileName := flag.Arg(0)
	input, err := os.Open(fileName)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer input.Close()

	content, err := io.ReadAll(input)
	if err != nil {
		fmt.Print(err)
		return
	}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		columns := strings.Split(string(line), ",")
		value1, err := strconv.Atoi(columns[0])
		if err != nil {
			panic(err)
		}
		value2, err := strconv.Atoi(columns[2])
		if err != nil {
			panic(err)
		}
		operator := columns[1]

		switch operator {
		case "+", "plus":
			fmt.Println(value1 + value2)
		case "-":
			fmt.Println(value1 - value2)
		case "*":
			fmt.Println(value1 * value2)
		case "/":
			fmt.Println(value1 / value2)
		default:
			fmt.Print("wrong operation")
		}
	}
}
