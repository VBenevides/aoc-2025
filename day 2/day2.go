package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func exampleInput() string {
	return "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
}

func realInput() string {
	return "197-407,262128-339499,557930-573266,25-57,92856246-93001520,2-12,1919108745-1919268183,48414903-48538379,38342224-38444598,483824-534754,1056-1771,4603696-4688732,75712519-75792205,20124-44038,714164-782292,4429019-4570680,9648251-9913729,6812551522-6812585188,58-134,881574-897488,648613-673853,5261723647-5261785283,60035-128980,9944818-10047126,857821365-857927915,206885-246173,1922-9652,424942-446151,408-1000"
}

func getRanges(input string) [][]int {
	ranges := [][]int{}
	parts := strings.Split(input, ",")
	for _, part := range parts {
		bounds := strings.Split(part, "-")
		val1, _ := strconv.Atoi(bounds[0])
		val2, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, []int{val1, val2})
	}
	return ranges
}

func part1(ranges [][]int) int {
	sum := 0
	for _, r := range ranges {
		for num := r[0]; num <= r[1]; num++ {
			if len(strconv.Itoa(num))%2 != 0 {
				continue
			}
			numStr := strconv.Itoa(num)
			mid := len(numStr) / 2
			firstHalf := numStr[:mid]
			secondHalf := numStr[mid:]
			if firstHalf == secondHalf {
				sum += num
			}
		}
	}
	return sum
}

func part2(ranges [][]int) int {
	sum := 0
	for _, r := range ranges {
		for num := r[0]; num <= r[1]; num++ {
			numStr := strconv.Itoa(num)
			numLen := len(numStr)
			for subSize := numLen / 2; subSize >= 1; subSize-- {
				if numLen%subSize != 0 {
					continue
				}
				baseSegment := numStr[:subSize]
				duplicated := true
				for i := subSize; i < numLen; i += subSize {
					if numStr[i:i+subSize] != baseSegment {
						duplicated = false
						break
					}
				}
				if duplicated {
					sum += num
					break
				}
			}
		}
	}
	return sum
}

func main() {
	t1 := time.Now()
	defer func() {
		t2 := time.Now()
		fmt.Printf("Execution time: %v\n", t2.Sub(t1))
	}()
	input := exampleInput()
	//input := realInput()
	ranges := getRanges(input)
	fmt.Printf("%v\n", ranges)
	part1 := part1(ranges)
	part2 := part2(ranges)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
