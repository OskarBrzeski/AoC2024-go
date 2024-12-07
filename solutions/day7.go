package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test7() {
	fmt.Println("Day 7 Test")
	fmt.Println("Part 1:", d7part1("./data/test7.txt"), "| Expected: 3749")
	fmt.Println("Part 2:", d7part2("./data/test7.txt"), "| Expected: 11387")
	fmt.Println()
}

func Day7() {
	fmt.Println("Day 7")
	fmt.Println("Part 1:", d7part1("./data/day7.txt"))
	fmt.Println("Part 2:", d7part2("./data/day7.txt"))
	fmt.Println()
}

func d7parseInput(filepath string) map[int][]int {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	trimmedFile := strings.TrimSpace(string(fileContent))
	lines := strings.Split(trimmedFile, "\n")

	result := map[int][]int{}
	for _, line := range lines {
		temp := strings.Split(line, ": ")
		total, _ := strconv.Atoi(temp[0])

		nums := []int{}
		for _, numstr := range strings.Split(temp[1], " ") {
			num, _ := strconv.Atoi(numstr)
			nums = append(nums, num)
		}
		result[total] = nums
	}

	return result
}

func d7part1(filename string) int {
	stuff := d7parseInput(filename)
	count := 0

	for total, nums := range stuff {
		if reduce_nums(total, nums[1:], nums[0], false) {
			count += total
		}
	}

	return count
}

func d7part2(filename string) int {
	stuff := d7parseInput(filename)
	count := 0

	for total, nums := range stuff {
		if reduce_nums(total, nums[1:], nums[0], true) {
			count += total
		}
	}

	return count
}

func reduce_nums(total int, nums []int, current int, extended bool) bool {
	if len(nums) < 1 {
		return current == total
	}

	return (reduce_nums(total, nums[1:], current+nums[0], extended) ||
		reduce_nums(total, nums[1:], current*nums[0], extended) ||
		(extended && reduce_nums(total, nums[1:], concat(current, nums[0]), extended)))
}

func concat(a, b int) int {
	result, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return result
}
