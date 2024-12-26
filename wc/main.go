package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		filePath := args[1]
		fmt.Printf("%v %v %v %s\n", GetLines(filePath), GetWords(filePath), GetBytes(filePath), filePath)
		os.Exit(0)
	}
	var action, filePath = args[1], args[2]
	switch action {
	case "-c":
		fmt.Printf("%s file is %v bytes\n", filePath, GetBytes(filePath))
	case "-l":
		fmt.Printf("%s file consists of %v lines\n", filePath, GetLines(filePath))
	case "-w":
		fmt.Printf("%s file has %v words\n", filePath, GetWords(filePath))
	case "-m":
		fmt.Printf("%s file has %v characters\n", filePath, GetCharacters(filePath))
	default:
		fmt.Printf("%s action not found\n", action)
	}
}
