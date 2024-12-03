package solutions

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func d3parseInput(filepath string) []string {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	sa := strings.Fields(string(fileContent))
	s := strings.Join(sa, "")

	re, err := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	if err != nil {
		fmt.Println("Regular Expression Did Not Compile")
		fmt.Println(err)
	}

	return re.FindAllString(s, -1)
}

func Test3() {
	fmt.Println("Day 3 Test")
	fmt.Println("Part 1:", d3part1("./data/test3p1.txt"), "| Expected: 161")
	fmt.Println("Part 2:", d3part2("./data/test3p2.txt"), "| Expected: 48")
	fmt.Println()
}

func Day3() {
	fmt.Println("Day 3")
	fmt.Println("Part 1:", d3part1("./data/day3.txt"))
	fmt.Println("Part 2:", d3part2("./data/day3.txt"))
	fmt.Println()
}

func d3part1(filename string) int {
	exprs := d3parseInput(filename)
	result := 0

	for _, expr := range exprs {
		if expr[0] != 'm' {
			continue
		}
		result += mul(expr)
	}

	return result
}

func d3part2(filename string) int {
	exprs := d3parseInput(filename)
	result := 0

	sum := true
	for _, expr := range exprs {
		if expr == "do()" {
			sum = true
			continue
		}

		if expr == "don't()" {
			sum = false
			continue
		}

		if sum {
			result += mul(expr)
		}
	}

	return result
}

func mul(expr string) int {
	nums := expr[4 : len(expr)-1]
	splits := strings.SplitN(nums, ",", 2)
	x, _ := strconv.Atoi(splits[0])
	y, _ := strconv.Atoi(splits[1])
	return x * y
}
