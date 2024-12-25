package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	var action, filePath = args[0], args[1]
	switch action {
	case "-c":
		// TODO: Step one completed
		fmt.Printf("%s file is %v bytes\n", filePath, getBytes(filePath))
	case "-l":
		// TODO: Step two
		fmt.Printf("%s file consists of %v lines\n", filePath, getLines(filePath))
	}

}

func getLines(file string) int {
	numberOfLines := 0
	data, err := os.ReadFile(file)
	check(err)
	for _, b := range data {
		if string(b) == "\n" {
			numberOfLines = numberOfLines + 1
		}
	}
	return numberOfLines
}

func getBytes(file string) int {
	data, err := os.ReadFile(file)
	check(err)
	return len(data)
}
