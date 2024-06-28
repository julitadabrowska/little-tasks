// the prohram is opening a file and printing the index of a header
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
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

	var value1index int
	var value2index int
	var operatorIndex int

	lines := strings.Split(string(content), "\n")
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

	fmt.Println(value1index)
	fmt.Println(value2index)
	fmt.Println(operatorIndex)
}
