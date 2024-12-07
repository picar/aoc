package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func atoislice(s string) []int {
	var r []int
	f := strings.Fields(s)
	for s := range f {
		i, _ := strconv.Atoi(f[s])
		r = append(r, i)
	}
	return r
}

type Order int

const (
	Ascending Order = iota
	Descending
	Unknown
)

func isSafe(s []int) bool {
	isAscending := Unknown

	var prev int
	for i := range s {
		if i == 0 {
			prev = s[i]
			continue
		} else {
			if abs(s[i]-prev) > 3 || s[i]-prev == 0 {
				return false
			}
			if s[i] > prev {
				if isAscending == Descending {
					return false
				} else {
					isAscending = Ascending
				}
			} else if s[i] < prev {
				if isAscending == Ascending {
					return false
				} else {
					isAscending = Descending
				}
			}
			prev = s[i]
		}
	}
	return true
}

func main() {
	file, err := os.Open("../../data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := atoislice(line)
		if isSafe(s) {
			count++
		}
	}
	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
