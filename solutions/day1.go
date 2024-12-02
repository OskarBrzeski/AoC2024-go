package solutions

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Test1() {
	fmt.Println("Day 1 Test")
	fmt.Println("Part 1:", d1part1("./data/test1.txt"), "| Expected: 11")
	fmt.Println("Part 2:", d1part2("./data/test1.txt"), "| Expected: 31")
	fmt.Println()
}

func Day1() {
	fmt.Println("Day 1")
	fmt.Println("Part 1:", d1part1("./data/day1.txt"))
	fmt.Println("Part 2:", d1part2("./data/day1.txt"))
	fmt.Println()
}

func d1part1(filename string) int {
	left, right := d1parseInput(filename)

	sort.Ints(left)
	sort.Ints(right)

	total := 0

	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			total += left[i] - right[i]
		} else {
			total += right[i] - left[i]
		}
	}

	return total
}

func d1part2(filename string) int {
	left, right := d1parseInput(filename)

	total := 0

	for i := 0; i < len(left); i++ {
		count := 0

		for j := 0; j < len(right); j++ {
			if right[j] == left[i] {
				count++
			}
		}

		total += left[i] * count
	}

	return total
}

func d1parseInput(filename string) ([]int, []int) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Fields(string(fileContent))

	var left []int
	var right []int

	for i := 0; i < len(lines)/2; i++ {
		numl, _ := strconv.Atoi(lines[2*i])
		left = append(left, numl)
		numr, _ := strconv.Atoi(lines[2*i+1])
		right = append(right, numr)
	}

	return left, right
}
