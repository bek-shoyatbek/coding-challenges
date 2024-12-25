package main

import "os"

func GetLines(file string) int {
	numberOfLines := 0
	data, err := os.ReadFile(file)
	CheckError(err)
	for _, b := range data {
		if string(b) == "\n" {
			numberOfLines = numberOfLines + 1
		}
	}
	return numberOfLines
}
