package main

import (
	"bufio"
	"fmt"
	"os"
)

func Getstring(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = file.Close()

	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func Getcount(lines []string) ([]int, []string) {
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	return counts, names
}

func main() {
	lines, _ := Getstring("../data/vote.txt")
	fmt.Println(lines)

	counts, names := Getcount(lines)
	for i, name := range names {
		fmt.Println(name, counts[i])
	}
}
