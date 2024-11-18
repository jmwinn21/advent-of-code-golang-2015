package main

import (
	"fmt"
	"os"

	"aoc/internal/utils"
)

func main() {
	input, err := utils.ReadFile("resources/day02.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	floor := 0

	for _, c := range input {
		if c == '(' {
			floor = floor + 1
		} else if c == ')' {
			floor = floor - 1
		}
	}

	return floor
}

// part two
func part2(input string) int {
	floor := 0
	basement := -1

	for i, c := range input {
		if c == '(' {
			floor = floor + 1
		} else if c == ')' {
			floor = floor - 1
		}

		if floor == -1 {
			basement = i + 1
			break
		}
	}

	return basement
}
