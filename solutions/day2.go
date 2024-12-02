package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test2() {
	fmt.Println("Day 2 Test")
	fmt.Println("Part 1:", d2part1("./data/test2.txt"), "| Expected: 2")
	fmt.Println("Part 2:", d2part2("./data/test2.txt"), "| Expected: 4")
	fmt.Println()
}

func Day2() {
	fmt.Println("Day 2")
	fmt.Println("Part 1:", d2part1("./data/day2.txt"))
	fmt.Println("Part 2:", d2part2("./data/day2.txt"))
	fmt.Println()
}

func d2part1(filepath string) int {
	reports := d2parseInput(filepath)
	total := 0

	for _, report := range reports {
		if d2reportSafe(report) {
			total++
		}
	}

	return total
}

func d2part2(filepath string) int {
	reports := d2parseInput(filepath)
	total := 0

	for _, report := range reports {
		for i := 0; i < len(report); i++ {
			var newReport []int
			newReport = append(newReport, report[:i]...)
			newReport = append(newReport, report[i+1:]...)

			if d2reportSafe(newReport) {
				total++
				break
			}
		}
	}

	return total
}

func d2reportSafe(report []int) bool {
	var increasing bool
	if report[0] < report[1] {
		increasing = true
	} else if report[0] > report[1] {
		increasing = false
	} else {
		return false
	}

	for i := 1; i < len(report); i++ {
		if report[i-1] == report[i] {
			return false
		}

		if report[i-1] < report[i] != increasing {
			return false
		}

		if report[i-1] > report[i] == increasing {
			return false
		}

		abs := report[i-1] - report[i]
		if abs < 0 {
			abs = -abs
		}
		if abs > 3 {
			return false
		}
	}

	return true
}

func d2parseInput(filepath string) [][]int {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	var lines []string = strings.Split(string(fileContent), "\n")

	var reports [][]int

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		var reportStr []string = strings.Split(lines[i], " ")
		var report []int

		for _, v := range reportStr {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			}

			report = append(report, num)
		}

		reports = append(reports, report)
	}

	return reports
}
