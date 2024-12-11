package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test10() {
	fmt.Println("Day 10 Test")
	fmt.Println("Part 1:", d10part1("./data/test10.txt"), "| Expected: 36")
	fmt.Println("Part 2:", d10part2("./data/test10.txt"), "| Expected: 81")
	fmt.Println()
}

func Day10() {
	fmt.Println("Day 10")
	fmt.Println("Part 1:", d10part1("./data/day10.txt"))
	fmt.Println("Part 2:", d10part2("./data/day10.txt"))
	fmt.Println()
}

func d10parseInput(filepath string) [][]int {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	trimmed := strings.TrimSpace(string(fileContent))
	lines := strings.Split(trimmed, "\n")

	result := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		result = append(result, row)
	}

	return result
}

func d10part1(filename string) int {
	topo_map := d10parseInput(filename)
	trailheads, trailends := find_trail(topo_map)

	count := 0
	for _, nine := range trailends {
		reach_map := create_reach_map(topo_map)
		fill_reach_map(topo_map, reach_map, nine.x, nine.y, 10)
		for _, zero := range trailheads {
			if reach_map[zero.y][zero.x] {
				count++
			}
		}
	}

	return count
}

func d10part2(filename string) int {
	topo_map := d10parseInput(filename)
	trailheads, trailends := find_trail(topo_map)

	count := 0
	for _, zero := range trailheads {
		score_map := create_score_map(topo_map)
		fill_score_map(topo_map, score_map, zero.x, zero.y, -1)
		for _, nine := range trailends {
			count += score_map[nine.y][nine.x]
		}
	}

	return count
}

type coords struct {
	x int
	y int
}

func find_trail(topo_map [][]int) ([]coords, []coords) {
	trailheads := []coords{}
	trailends := []coords{}
	for y, line := range topo_map {
		for x, num := range line {
			if num == 0 {
				trailheads = append(trailheads, coords{x, y})
			}

			if num == 9 {
				trailends = append(trailends, coords{x, y})
			}
		}
	}
	return trailheads, trailends
}

func create_reach_map(topo_map [][]int) [][]bool {
	reach_map := make([][]bool, len(topo_map))
	for i, line := range topo_map {
		reach_map[i] = make([]bool, len(line))
	}

	return reach_map
}

func fill_reach_map(topo_map [][]int, reach_map [][]bool, x, y int, prev int) {
	if reach_map[y][x] {
		return
	}

	if prev == 0 {
		return
	}

	current := topo_map[y][x]

	if prev-current == 1 {
		reach_map[y][x] = true

		bound_y := len(topo_map)
		bound_x := len(topo_map[0])
		if x > 0 {
			fill_reach_map(topo_map, reach_map, x-1, y, current)
		}
		if x < bound_x-1 {
			fill_reach_map(topo_map, reach_map, x+1, y, current)
		}
		if y > 0 {
			fill_reach_map(topo_map, reach_map, x, y-1, current)
		}
		if y < bound_y-1 {
			fill_reach_map(topo_map, reach_map, x, y+1, current)
		}
	}
}

func create_score_map(topo_map [][]int) [][]int {
	score_map := make([][]int, len(topo_map))
	for i, line := range topo_map {
		score_map[i] = make([]int, len(line))
	}

	return score_map
}

func fill_score_map(topo_map [][]int, score_map [][]int, x, y int, prev int) {
	if prev == 9 {
		return
	}

	current := topo_map[y][x]

	if current-prev == 1 {
		score_map[y][x]++

		bound_y := len(topo_map)
		bound_x := len(topo_map[0])
		if x > 0 {
			fill_score_map(topo_map, score_map, x-1, y, current)
		}
		if x < bound_x-1 {
			fill_score_map(topo_map, score_map, x+1, y, current)
		}
		if y > 0 {
			fill_score_map(topo_map, score_map, x, y-1, current)
		}
		if y < bound_y-1 {
			fill_score_map(topo_map, score_map, x, y+1, current)
		}
	}
}
