package day01

import (
	"fmt"
	"sort"

	aocutil "aoc/util"
)

func Puzzle2() {
	fmt.Println("Day 1 / Puzzle 2!")

	l1, l2 := aocutil.ReadInput("input.txt")

	sort.Ints(l1)
	sort.Ints(l2)

	total := 0

	for i := 0; i < len(l1); i++ {
		current := l1[i]
		mult := 0
		for j := 0; j < len(l2); j++ {
			if l2[j] < current {
				continue
			}
			if l2[j] > current {
				break
			}
			mult += 1
		}
		total += (current * mult)
	}

	fmt.Printf("\tAnswer = %d\n", total)
}
