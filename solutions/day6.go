package solutions

import (
	"fmt"
	"os"
	"strings"
)

func Test6() {
	fmt.Println("Day 6 Test")
	fmt.Println("Part 1:", d6part1("./data/test6.txt"), "| Expected: 41")
	fmt.Println("Part 2:", d6part2("./data/test6.txt"), "| Expected: 6")
	fmt.Println()
}

func Day6() {
	fmt.Println("Day 6")
	fmt.Println("Part 1:", d6part1("./data/day6.txt"))
	fmt.Println("Part 2:", d6part2("./data/day6.txt"))
	fmt.Println()
}

func d6parseInput(filepath string) [][]rune {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	grid := [][]rune{}
	for _, line := range strings.Fields(string(fileContent)) {
		runes := []rune{}
		for _, char := range line {
			runes = append(runes, char)
		}
		grid = append(grid, runes)
	}

	return grid
}

const (
	UP    = iota
	RIGHT = iota
	DOWN  = iota
	LEFT  = iota
)

type guardtype struct {
	i   int
	j   int
	dir int
}

func d6part1(filename string) int {
	grid := d6parseInput(filename)
	i, j := locate_guard(grid)
	guard := guardtype{i, j, UP}
	for !guard_outside(guard, grid) {
		guard = move_guard(guard, grid)
	}
	return count_X(grid)
}

func d6part2(filename string) int {
	grid := d6parseInput(filename)
	start_i, start_j := locate_guard(grid)

	count := 0

	for i, line := range grid {
		for j, char := range line {
			if char != '.' {
				continue
			}

			grid_copy := grid_deepcopy(grid)
			grid_copy[i][j] = '#'

			previous := map[guardtype]bool{}
			guard := guardtype{start_i, start_j, UP}
			for !guard_outside(guard, grid) {
				if previous[guard] {
					count++
					break
				}
				previous[guard] = true
				guard = move_guard(guard, grid_copy)
			}
		}
	}

	return count
}

func grid_deepcopy(grid [][]rune) [][]rune {
	grid_copy := [][]rune{}

	for _, line := range grid {
		line_copy := []rune{}
		for _, char := range line {
			line_copy = append(line_copy, char)
		}
		grid_copy = append(grid_copy, line_copy)
	}
	return grid_copy
}

func locate_guard(grid [][]rune) (i, j int) {
	for i, line := range grid {
		for j, char := range line {
			if char == '^' {
				return i, j
			}
		}
	}

	return -1, -1
}

func guard_outside(guard guardtype, grid [][]rune) bool {
	if guard.i < 0 {
		return true
	}
	if guard.i >= len(grid) {
		return true
	}
	if guard.j < 0 {
		return true
	}
	if guard.j >= len(grid[0]) {
		return true
	}
	return false
}

func move_guard(guard guardtype, grid [][]rune) guardtype {
	switch guard.dir {
	case UP:
		if guard.i-1 < 0 {
			grid[guard.i][guard.j] = 'X'
			guard.i--
			break
		}
		if grid[guard.i-1][guard.j] == '#' {
			grid[guard.i][guard.j] = '>'
			guard.dir = RIGHT
		} else {
			grid[guard.i][guard.j] = 'X'
			grid[guard.i-1][guard.j] = '^'
			guard.i--
		}
	case RIGHT:
		if guard.j+1 >= len(grid[0]) {
			grid[guard.i][guard.j] = 'X'
			guard.j++
			break
		}
		if grid[guard.i][guard.j+1] == '#' {
			guard.dir = DOWN
			grid[guard.i][guard.j] = 'v'
		} else {
			grid[guard.i][guard.j] = 'X'
			grid[guard.i][guard.j+1] = '>'
			guard.j++
		}
	case DOWN:
		if guard.i+1 >= len(grid) {
			grid[guard.i][guard.j] = 'X'
			guard.i++
			break
		}
		if grid[guard.i+1][guard.j] == '#' {
			grid[guard.i][guard.j] = '<'
			guard.dir = LEFT
		} else {
			grid[guard.i][guard.j] = 'X'
			grid[guard.i+1][guard.j] = 'v'
			guard.i++
		}
	case LEFT:
		if guard.j-1 < 0 {
			grid[guard.i][guard.j] = 'X'
			guard.j--
			break
		}
		if grid[guard.i][guard.j-1] == '#' {
			grid[guard.i][guard.j] = '^'
			guard.dir = UP
		} else {
			grid[guard.i][guard.j] = 'X'
			grid[guard.i][guard.j-1] = '<'
			guard.j--
		}
	}

	return guard
}

func count_X(grid [][]rune) int {
	count := 0
	for _, line := range grid {
		for _, char := range line {
			if char == 'X' {
				count++
			}
		}
	}
	return count
}

func _print_grid(grid [][]rune) {
	for _, line := range grid {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
