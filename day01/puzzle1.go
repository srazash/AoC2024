package day01

import (
	"fmt"

	aocutil "aoc/util"
)

func Puzzle1() {
	fmt.Println("Day 1 / Puzzle 1!")

	l1, l2 := aocutil.ReadInput("input.txt")

	fmt.Printf("l1 = %d, l2 = %d\n", len(*l1), len(*l2))
}
