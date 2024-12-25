package main

import (
	"bufio"
	"os"
)

func GetWords(filePath string) int {
	file, err := os.Open(filePath)
	CheckError(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)

	count := 0

	for fileScanner.Scan() {
		count++
	}

	err = fileScanner.Err()
	CheckError(err)
	return count

}
