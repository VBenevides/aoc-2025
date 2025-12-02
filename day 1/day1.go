package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func exampleInput() []string {
	return []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
}

func readInput(filename string) []string {
	content, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	return lines
}

func signal(x string) int {
	if strings.Contains(x, "L") {
		return -1
	}
	return 1
}

func value(x string) int {
	val, _ := strconv.Atoi(x[1:])
	return signal(x) * val
}

func part1(lines []string) int {
	counter := 0
	sum := 50
	for _, x := range lines {
		sum += value(x)
		sum %= 100
		if sum == 0 {
			counter++
		}
	}
	return counter
}

func part2(lines []string) int {
	counter := 0
	sum := 50
	for _, x := range lines {
		// Wrap-around logic
		if sum == 100 && signal(x) == 1 {
			sum = 0
		}
		if sum == 0 && signal(x) == -1 {
			sum = 100
		}
		sum += value(x)

		for {
			if sum >= 0 {
				break
			}
			sum += 100
			counter += 1
		}

		for {
			if sum <= 100 {
				break
			}
			sum -= 100
			counter += 1
		}
		if sum == 100 {
			sum = 0
		}
		if sum == 0 {
			counter++
		}
	}
	return counter
}

func main() {
	//input := exampleInput()
	input := readInput("input.txt")
	result := part1(input)
	fmt.Println("Part 1:", result)
	result = part2(input)
	fmt.Println("Part 2:", result)
}
