package solutions

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	fmt.Println("Day 1 Part 1:", d1part1())
	fmt.Println("Day 1 Part 2:", d1part2())
}

func d1part1() int {
	left, right := d1parseInput()

	sort.Ints(left[:])
	sort.Ints(right[:])

	var total int = 0

	for i := 0; i < 1000; i++ {
		if left[i] > right[i] {
			total += left[i] - right[i]
		} else {
			total += right[i] - left[i]
		}
	}

	return total
}

func d1part2() int {
	left, right := d1parseInput()

	total := 0

	for i := 0; i < 1000; i++ {
		count := 0

		for j := 0; j < 1000; j++ {
			if right[j] == left[i] {
				count++
			}
		}

		total += left[i] * count
	}

	return total
}

func d1parseInput() ([1000]int, [1000]int) {
	fileContent, err := os.ReadFile("./data/day1.txt")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Fields(string(fileContent))

	var left [1000]int
	var right [1000]int

	for i := 0; i < 1000; i++ {
		left[i], _ = strconv.Atoi(lines[2*i])
		right[i], _ = strconv.Atoi(lines[2*i+1])
	}

	return left, right
}
