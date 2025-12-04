package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func exampleInput() [][]rune {
	x := `..@@.@@@@.
        @@@.@.@.@@
        @@@@@.@.@@
        @.@@@@..@.
        @@.@@@@.@@
        .@@@@@@@.@
        .@.@.@.@@@
        @.@@@.@@@@
        .@@@@@@@@.
        @.@.@@@.@.`
	stringSlice := strings.Split(strings.TrimSpace(string(x)), "\n")
	input := make([][]rune, len(stringSlice))
	for i, line := range stringSlice {
		input[i] = []rune(strings.TrimSpace(string(line)))
	}
	return input
}

func realInput() [][]rune {
	content, _ := os.ReadFile("input.txt")
	stringSlice := strings.Split(strings.TrimSpace(string(content)), "\n")
	input := make([][]rune, len(stringSlice))
	for i, line := range stringSlice {
		input[i] = []rune(strings.TrimSpace(string(line)))
	}
	return input
}

func checkAccessible(rolls [][]rune, row, col, limit int) bool {
	numRows := len(rolls)
	numCols := len(rolls[0])

	adjacentCount := 0

	if row > 0 {
		if col > 0 && rolls[row-1][col-1] == '@' {
			adjacentCount++
		}
		if rolls[row-1][col] == '@' {
			adjacentCount++
		}
		if col < numCols-1 && rolls[row-1][col+1] == '@' {
			adjacentCount++
		}
	}

	if col > 0 && rolls[row][col-1] == '@' {
		adjacentCount++
	}
	if col < numCols-1 && rolls[row][col+1] == '@' {
		adjacentCount++
	}

	if row < numRows-1 {
		if col > 0 && rolls[row+1][col-1] == '@' {
			adjacentCount++
		}
		if rolls[row+1][col] == '@' {
			adjacentCount++
		}
		if col < numCols-1 && rolls[row+1][col+1] == '@' {
			adjacentCount++
		}
	}

	return adjacentCount <= limit
}

func part1(rolls [][]rune) {
	accessibleCount := 0
	for i, row := range rolls {
		for j, value := range row {
			if value == '@' && checkAccessible(rolls, i, j, 3) {
				accessibleCount++
			}
		}
	}
	println("Part 1:", accessibleCount)
}

func removeRoll(rolls [][]rune, row, col int) {
	rolls[row][col] = 'x'
}

func part2(rolls [][]rune) {
	accessibleCount := 0
	for k := 0; k < len(rolls); k++ {
		for i, row := range rolls {
			for j, value := range row {
				if value == '@' && checkAccessible(rolls, i, j, 3) {
					accessibleCount++
					removeRoll(rolls, i, j)
				}
			}
		}
	}
	println("Part 2:", accessibleCount)
}

func main() {
	t1 := time.Now()
	defer func() {
		fmt.Println("Execution Time: ", time.Since(t1))
	}()
	rolls := realInput()
	part1(rolls)
	part2(rolls)
}
