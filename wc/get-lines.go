package main

import (
	"bufio"
	"os"
)

func GetLines(filePath string) int {
	numberOfLines := 0
	file, err := os.Open(filePath)
	CheckError(err)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		numberOfLines = numberOfLines + 1
	}
	return numberOfLines
}
