package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test5() {
	fmt.Println("Day 5 Test")
	fmt.Println("Part 1:", d5part1("./data/test5.txt"), "| Expected: 143")
	fmt.Println("Part 2:", d5part2("./data/test5.txt"), "| Expected: 123")
	fmt.Println()
}

func Day5() {
	fmt.Println("Day 5")
	fmt.Println("Part 1:", d5part1("./data/day5.txt"))
	fmt.Println("Part 2:", d5part2("./data/day5.txt"))
	fmt.Println()
}

func d5parseInput(filepath string) (map[int][]int, [][]int) {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	temp := strings.Split(string(fileContent), "\n\n")

	return parse_rules(temp[0]), parse_updates(temp[1])
}

func parse_rules(content string) map[int][]int {
	rules := make(map[int][]int)
	for _, rule := range strings.Split(content, "\n") {
		both := strings.Split(rule, "|")
		before, _ := strconv.Atoi(both[0])
		after, _ := strconv.Atoi(both[1])
		if rules[before] != nil {
			rules[before] = append(rules[before], after)
		} else {
			rules[before] = []int{after}
		}
	}
	return rules
}

func parse_updates(content string) [][]int {
	updates := [][]int{}
	for _, update := range strings.Split(content, "\n") {
		nums := strings.Split(update, ",")
		u_arr := []int{}
		for _, item := range nums {
			u_int, _ := strconv.Atoi(item)
			u_arr = append(u_arr, u_int)
		}
		updates = append(updates, u_arr)
	}
	return updates
}

func d5part1(filename string) int {
	rules, updates := d5parseInput(filename)
	count := 0

	for _, update := range updates {
		if is_update_correct(rules, update) {
			count += update[(len(update)-1)/2]
		}
	}

	return count
}

func d5part2(filename string) int {
	rules, updates := d5parseInput(filename)
	count := 0

	for _, update := range updates {
		if !is_update_correct(rules, update) {
			count += fixed_middle(rules, update)
		}
	}

	return count
}

func is_update_correct(rules map[int][]int, update []int) bool {
	for i := 1; i < len(update); i++ {
		before := rules[update[i]]

		for _, b := range before {
			if b == update[i-1] {
				return false
			}
		}
	}
	return true
}

func fixed_middle(rules map[int][]int, update []int) int {
	for j := len(update); j > 0; j-- {
		for i := 1; i < j; i++ {
			before := rules[update[i]]
			for _, b := range before {
				if b == update[i-1] {
					update[i-1], update[i] = update[i], update[i-1]
				}
			}
		}
	}

	return update[(len(update)-1)/2]
}
