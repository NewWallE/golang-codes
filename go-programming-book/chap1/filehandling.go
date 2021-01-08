package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for _, filename := range os.Args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println(fmt.Errorf("some issue while opening file%s", filename))
			return
		}
		fmt.Println(findDuplicates(f))
	}
}

func findDuplicates(file *os.File) int {
	input := bufio.NewScanner(file)
	lineMap := make(map[string]int)
	for input.Scan() {
		lineMap[input.Text()]++
	}
	return countDuplicates(lineMap)
}

func countDuplicates(lineMap map[string]int) int {
	duplicates := 0
	for line, count := range lineMap {
		if count > 1 {
			fmt.Printf("Found duplicate line: %s\n", line)
			duplicates++
		}
	}
	return duplicates
}
