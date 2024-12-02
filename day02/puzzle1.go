package day02

import (
	"fmt"

	aocutil "aoc/util"
)

func Puzzle1() {
	fmt.Println("Day 2 / Puzzle 1!")

	l := aocutil.Day02Input("day02/input.txt")

	total := 0

	for _, v := range l {
		if safe(v) {
			total += 1
		}
	}

	fmt.Printf("\tAnswer = %d\n", total)
}

func safe(input []int) bool {
	stepUp := input[0] < input[1]

	for i := range input {
		if i+1 == len(input) {
			break
		}
		if stepUp {
			if input[i] > input[i+1] {
				return false
			}
			diff := input[i+1] - input[i]
			if diff < 1 || diff > 3 {
				return false
			}
		} else {
			if input[i] < input[i+1] {
				return false
			}
			diff := input[i] - input[i+1]
			if diff < 1 || diff > 3 {
				return false
			}
		}
	}
	return true
}
