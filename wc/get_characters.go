package main

import (
	"bufio"
	"os"
)

func GetCharacters(filePath string) int {
	file, err := os.Open(filePath)
	CheckError(err)

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanRunes)

	count := 0

	for fileScanner.Scan() {
		count = count + 1
	}

	return count
}
