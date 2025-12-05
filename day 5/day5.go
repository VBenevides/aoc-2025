package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func exampleInput() []string {
	input := `3-5
		10-14
		16-20
		12-18

		1
		5
		8
		11
		17
		32
	`
	return strings.Split(strings.TrimSpace(input), "\n")
}

func realInput() []string {
	content, _ := os.ReadFile("input.txt")
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func checkInRanges(ranges [][]int, num int) bool {
	for _, r := range ranges {
		if num >= r[0] && num <= r[1] {
			return true
		}
	}
	return false
}

func listIngredients(database []string) ([][]int, []int) {
	freshRanges := [][]int{}
	ingredientList := []int{}
	for _, line := range database {
		if !strings.Contains(line, "-") {
			num, _ := strconv.Atoi(strings.TrimSpace(line))
			ingredientList = append(ingredientList, num)
			continue
		}

		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

		freshRanges = append(freshRanges, []int{start, end})
	}

	return freshRanges, ingredientList
}

func part1(freshRanges [][]int, ingredientList []int) {
	count := 0
	for _, ingredient := range ingredientList {
		if checkInRanges(freshRanges, ingredient) {
			count += 1
		}
	}
	println("Part 1:", count)
}

func part2(freshRanges [][]int) {
	numRanges := len(freshRanges)
	starts := make([]int, numRanges)
	ends := make([]int, numRanges)
	for i, r := range freshRanges {
		starts[i] = r[0]
		ends[i] = r[1]
	}

	slices.Sort(starts)
	slices.Sort(ends)

	count := 0
	for i := range numRanges {
		if i == 0 {
			count += ends[0] - starts[0] + 1
			continue
		}

		if starts[i] <= ends[i-1] {
			starts[i] = ends[i-1] + 1
		}

		count += ends[i] - starts[i] + 1
	}
	println("Part 2:", count)
}

func main() {
	t1 := time.Now()
	defer func() {
		fmt.Println("Execution Time:", time.Since(t1))
	}()
	database := realInput()
	freshRanges, ingredientList := listIngredients(database)
	part1(freshRanges, ingredientList)
	part2(freshRanges)

}
