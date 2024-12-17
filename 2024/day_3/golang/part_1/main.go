package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func main() {
	file, err := os.Open("../../data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`)

	scanner := bufio.NewScanner(file)
	result := 0
	do := true
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
        for _, match := range matches {
			if match[0] == "do()" {
				// fmt.Println("do")
				do = true
			} else if match[0] == "don't()" {
				// fmt.Println("don't")
				do = false
			}
			if strings.HasPrefix(match[0], "mul") && do {
				// fmt.Println("mul")
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				result += num1 * num2
			}
        }
	}

	fmt.Println("Result:", result)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
