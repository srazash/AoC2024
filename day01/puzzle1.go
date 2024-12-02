package day01

import (
	"fmt"
	"sort"

	aocutil "aoc/util"
)

func Puzzle1() {
	fmt.Println("Day 1 / Puzzle 1!")

	l1, l2 := aocutil.ReadInput("input.txt")

	sort.Ints(l1)
	sort.Ints(l2)

	total := 0

	for i := 0; i < len(l1); i++ {
		if l1[i] > l2[i] {
			total += (l1[i] - l2[i])
		} else {
			total += (l2[i] - l1[i])
		}
	}

	fmt.Printf("\tAnswer = %d\n", total)
}
