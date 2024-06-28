// the program takes the csv, makes the calculation and returns the values and the result of the calculation in a new csv. It also recognizes the headers and returns the outcome in a proper order
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

	content, err := io.ReadAll(input)
	if err != nil {
		fmt.Print(err)
		return
	}

	defer input.Close()

	outputCsv := "outputFile.csv"

	output, err := os.Create(outputCsv)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer output.Close()

	var value1index int
	var value2index int
	var operatorIndex int

	lines := strings.Split(string(content), "\n")

	header := lines[0]
	output.WriteString(header + "\n")

	headers := strings.Split(lines[0], ",")
	if "value1" == headers[0] {
		value1index = 0
	} else if "value1" == headers[1] {
		value1index = 1
	} else if "value1" == headers[2] {
		value1index = 2
	}
	if "value2" == headers[0] {
		value2index = 0
	} else if "value2" == headers[1] {
		value2index = 1
	} else if "value2" == headers[2] {
		value2index = 2
	}
	if "operator" == headers[0] {
		operatorIndex = 0
	} else if "operator" == headers[1] {
		operatorIndex = 1
	} else if "operator" == headers[2] {
		operatorIndex = 2
	}

	for i, line := range lines {
		if i == 0 {
			continue
		}

		columns := strings.Split(string(line), ",")

		value1, err := strconv.Atoi(columns[value1index])
		if err != nil {
			panic(err)
		}
		value2, err := strconv.Atoi(columns[value2index])
		if err != nil {
			panic(err)
		}
		operator := columns[operatorIndex]

		if value1index == 0 && operatorIndex == 1 && value2index == 2 {
			switch operator {
			case "+", "plus":
				fmt.Println(value1 + value2)
				//fmt.Fprintln(output, value1+value2)
				fmt.Fprintf(output, "%d,%s,%d,%d\n", value1, operator, value2, value1+value2)
			case "-":
				fmt.Println(value1 - value2)
				//fmt.Fprintln(output, value1-value2)
				fmt.Fprintf(output, "%d,%s,%d,%d\n", value1, operator, value2, value1-value2)
			case "*":
				fmt.Println(value1 * value2)
				//fmt.Fprintln(output, value1*value2)
				fmt.Fprintf(output, "%d,%s,%d,%d\n", value1, operator, value2, value1*value2)
			case "/":
				fmt.Println(value1 / value2)
				//fmt.Fprintln(output, value1/value2)
				fmt.Fprintf(output, "%d,%s,%d,%d\n", value1, operator, value2, value1/value2)
			default:
				fmt.Print("wrong operation")
			}
		} else if value1index == 0 && operatorIndex == 2 && value2index == 1 {
			switch operator {
			case "+", "plus":
				fmt.Println(value1 + value2)
				//fmt.Fprintln(output, value1+value2)
				fmt.Fprintf(output, "%d,%d,%s,%d\n", value1, value2, operator, value1+value2)
			case "-":
				fmt.Println(value1 - value2)
				//fmt.Fprintln(output, value1-value2)
				fmt.Fprintf(output, "%d,%d,%s,%d\n", value1, value2, operator, value1-value2)
			case "*":
				fmt.Println(value1 * value2)
				//fmt.Fprintln(output, value1*value2)
				fmt.Fprintf(output, "%d,%d,%s,%d\n", value1, value2, operator, value1*value2)
			case "/":
				fmt.Println(value1 / value2)
				//fmt.Fprintln(output, value1/value2)
				fmt.Fprintf(output, "%d,%d,%s,%d\n", value1, value2, operator, value1/value2)
			default:
				fmt.Print("wrong operation")
			}
		}
	}
}

