package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func exampleInput() string {
	content, _ := os.ReadFile("example.txt")
	return string(content)
}

func realInput() string {
	content, _ := os.ReadFile("input.txt")
	return string(content)
}

func convertStringsToRuneMatrix(ss []string) [][]rune {
	runeMatrix := make([][]rune, len(ss))

	for i, s := range ss {
		runeMatrix[i] = []rune(s)
	}

	return runeMatrix
}

func transposeRuneMatrix(matrix [][]rune) [][]rune {
	if len(matrix) == 0 {
		return [][]rune{}
	}

	numRows := len(matrix)
	numCols := len(matrix[0])
	transposed := make([][]rune, numCols)

	for c := 0; c < numCols; c++ {
		transposed[c] = make([]rune, numRows)
		for r := 0; r < numRows; r++ {
			transposed[c][r] = matrix[r][c]
		}
	}

	return transposed
}

func prepareInput(input string, readVertical bool) ([][]int, []string) {
	lines := strings.Split(input, "\n")

	// Get operations
	var operations []string
	for _, x := range strings.Split(lines[len(lines)-1], " ") {
		if x != "" {
			operations = append(operations, strings.TrimSpace(x))
		}
	}

	// Get numbers
	runes := [][]rune{}
	runeLines := lines[:len(lines)-1]
	for i, lineStr := range runeLines {
		runes = append(runes, []rune{})
		for _, r := range lineStr {
			runes[i] = append(runes[i], r)
		}
	}

	// Get operation separators
	seps := []int{}
	numRows := len(runeLines)
	numCols := len(runeLines[0])
	for col := 0; col < numCols; col++ {
		onlySpace := true
		for row := 0; row < numRows; row++ {
			if runeLines[row][col] != ' ' {
				onlySpace = false
				break
			}
		}

		if onlySpace {
			seps = append(seps, col)
		}
	}

	// Split numbers based on separators
	seps = append(seps, numCols)
	rawNumbers := [][]string{}
	prevSep := -1
	for _, sep := range seps {
		group := []string{}
		for row := 0; row < numRows; row++ {
			part := string(runeLines[row][prevSep+1 : sep])
			group = append(group, part)
		}
		rawNumbers = append(rawNumbers, group)
		prevSep = sep
	}

	// Read depending on flag
	numbers := [][]int{}
	if readVertical {
		for _, lineStr := range rawNumbers {
			numRow := []int{}
			runeMatrix := convertStringsToRuneMatrix(lineStr)
			transposed := transposeRuneMatrix(runeMatrix)
			for _, runeSlice := range transposed {
				numStr := string(runeSlice)
				x, _ := strconv.Atoi(strings.TrimSpace(numStr))
				numRow = append(numRow, x)
			}
			numbers = append(numbers, numRow)
		}
	} else {
		for _, lineStr := range rawNumbers {
			numRow := []int{}
			for _, compStr := range lineStr {
				x, _ := strconv.Atoi(strings.TrimSpace(compStr))
				numRow = append(numRow, x)
			}
			numbers = append(numbers, numRow)
		}
	}
	return numbers, operations
}

func calculate(numbers [][]int, operations []string) {
	sum := 0
	for colId, op := range operations {
		var tmp int
		if op == "*" {
			tmp = 1
		} else {
			tmp = 0
		}

		groupNumbers := numbers[colId]
		for _, val := range groupNumbers {
			if op == "*" {
				tmp *= val
			} else {
				tmp += val
			}
		}
		sum += tmp
	}
	fmt.Printf("Result: %d\n", sum)
}

func main() {
	t1 := time.Now()
	defer func() {
		fmt.Printf("Execution Time: %s\n", time.Since(t1))
	}()
	input := realInput()
	numbers, operations := prepareInput(input, false)
	calculate(numbers, operations)

	numbers, operations = prepareInput(input, true)
	calculate(numbers, operations)
}
