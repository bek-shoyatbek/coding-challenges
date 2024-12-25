package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	var action, filePath = args[0], args[1]
	switch action {
	case "-c":
		// TODO: Step one completed
		fmt.Printf("%s file is %v bytes\n", filePath, GetBytes(filePath))
	case "-l":
		// TODO: Step two
		fmt.Printf("%s file consists of %v lines\n", filePath, GetLines(filePath))
	case "-w":
		// TODO: Step three
		fmt.Printf("%s file has %v words\n", filePath, GetWords(filePath))
	default:
		fmt.Printf("%s command not found", action)
	}
}
