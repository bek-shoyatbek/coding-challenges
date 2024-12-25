package main

import "os"

func GetBytes(file string) int {
	data, err := os.ReadFile(file)
	CheckError(err)
	return len(data)
}
