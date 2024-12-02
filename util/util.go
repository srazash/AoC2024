package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day01Input(f string) ([]int, []int) {
	var l1, l2 []int

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		v := strings.Split(scanner.Text(), "   ")
		v1, err := strconv.Atoi(v[0])
		if err != nil {
			log.Fatal(err)
		}
		v2, err := strconv.Atoi(v[1])
		if err != nil {
			log.Fatal(err)
		}

		l1 = append(l1, v1)
		l2 = append(l2, v2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(l1) != len(l2) {
		log.Fatal("list lengths do not match!")
	}

	return l1, l2
}

func Day02Input(f string) [][]int {
	var l [][]int

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ln := strings.Split(scanner.Text(), " ")
		var s []int

		for _, v := range ln {
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

			s = append(s, i)
		}

		l = append(l, s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return l
}
