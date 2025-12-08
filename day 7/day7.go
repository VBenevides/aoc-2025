package main

import (
	"fmt"
	"os"
	"strings"
)

var globalCounter int

func exampleInput() string {
	content, _ := os.ReadFile("example.txt")
	return strings.Replace(string(content), "S", "|", -1)

}

func realInput() string {
	content, _ := os.ReadFile("input.txt")
	return strings.Replace(string(content), "S", "|", -1)
}

func printMatrix(matrix [][]rune) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%c", val)
		}
		fmt.Println()
	}
}

func prepareInput(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	runeMatrix := make([][]rune, len(lines))

	for i, s := range lines {
		runeMatrix[i] = []rune(s)
	}

	return runeMatrix
}

func propagatePart1(matrix [][]rune) ([][]rune, int) {
	counter := 0
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i-1][j] != '|' {
				continue
			}
			if matrix[i][j] == '.' {
				matrix[i][j] = '|'
			}
			if matrix[i][j] == '^' {
				counter += 1
				if j > 0 {
					matrix[i][j-1] = '|'
				}
				if j < len(matrix[i])-1 {
					matrix[i][j+1] = '|'
				}
			}
		}
	}

	return matrix, counter
}

func propagatePart2(matrix [][]rune) ([][]rune, int) {
	counter := 0
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i-1][j] != '|' {
				continue
			}
			if matrix[i][j] == '.' {
				matrix[i][j] = '|'
			}
			if matrix[i][j] == '^' {
				counter += 1
				if j > 0 {
					matrix[i][j-1] = '|'
				}
				if j < len(matrix[i])-1 {
					matrix[i][j+1] = '|'
				}
			}
		}
	}

	return matrix, counter
}

func main() {
	input := exampleInput()
	matrix := prepareInput(input)
	matrix, counter := propagatePart1(matrix)
	fmt.Printf("Counter PT1: %d\n", counter)
	matrix = prepareInput(input)
	matrix, counter = propagatePart2(matrix)
	fmt.Printf("Counter PT2: %d\n", counter)
}
