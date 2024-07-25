// the program is opening a csv file, splitting it to rows and columns and returning a certain column/row/cell from the file
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Row struct {
	Id    int
	Count int
	Value int
}

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
	var data []Row
	for _, line := range lines {
		columns := strings.Split(string(line), ",")
		id, err := strconv.Atoi(columns[0])
		if err != nil {
			panic(err)
		}
		count, err := strconv.Atoi(columns[1])
		if err != nil {
			panic(err)
		}
		value, err := strconv.Atoi(columns[2])
		if err != nil {
			panic(err)
		}
		d := Row{
			Id:    id,
			Count: count,
			Value: value,
		}
		data = append(data, d)
		fmt.Println(d.Value)
	}
}
