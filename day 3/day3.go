package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func exampleInput() []string {
	return []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
}

func realInput() []string {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	return lines
}

func joltage(line string, numElements int) int {
	startIndex := -1
	value := 0
	for rightBlock := -numElements + 1; rightBlock <= 0; rightBlock++ {
		_startIndex := startIndex + 1

		subLine := line[_startIndex : len(line)+rightBlock]
		var maxRune rune
		var maxIndex int
		for i, r := range subLine {
			if i == 0 || r > maxRune {
				maxRune = r
				maxIndex = i
			}
		}
		value = value*10 + int(maxRune-'0')
		startIndex += maxIndex + 1
	}
	return value
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += joltage(line, 2)
	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += joltage(line, 12)
	}
	return sum
}

func main() {
	t1 := time.Now()
	defer func() {
		fmt.Println("Execution Time:", time.Since(t1))
	}()
	lines := realInput()
	println("Part 1:", part1(lines))
	println("Part 2:", part2(lines))
}
