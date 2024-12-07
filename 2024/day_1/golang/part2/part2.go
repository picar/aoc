package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	file, err := os.Open("../../data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var left, right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		i, _ := strconv.Atoi(f[0])
		left = append(left, i)
		i, _ = strconv.Atoi(f[1])
		right = append(right, i)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	right_count := make(map[int]int)

	for i := range right {
		right_count[right[i]]++
	}

	s := 0
	for i := range left {
		s += abs(left[i] * right_count[left[i]])
	}
	fmt.Println(s)
}