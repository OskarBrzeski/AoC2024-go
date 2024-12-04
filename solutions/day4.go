package solutions

import (
	"fmt"
	"os"
	"strings"
)

func Test4() {
	fmt.Println("Day 4 Test")
	fmt.Println("Part 1:", d4part1("./data/test4.txt"), "| Expected: 18")
	fmt.Println("Part 2:", d4part2("./data/test4.txt"), "| Expected: 9")
	fmt.Println()
}

func Day4() {
	fmt.Println("Day 4")
	fmt.Println("Part 1:", d4part1("./data/day4.txt"))
	fmt.Println("Part 2:", d4part2("./data/day4.txt"))
	fmt.Println()
}

func d4parseInput(filepath string) []string {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	return strings.Fields(string(fileContent))
}

func d4part1(filename string) int {
	lines := d4parseInput(filename)
	total := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if horizontal(i, j, lines) {
				total++
			}
			if vertical(i, j, lines) {
				total++
			}
			if diagonal_right(i, j, lines) {
				total++
			}
			if diagonal_left(i, j, lines) {
				total++
			}
		}
	}
	return total
}

func d4part2(filename string) int {
	lines := d4parseInput(filename)
	total := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if cross_mas(i, j, lines) {
				total++
			}
		}
	}

	return total
}

func horizontal(i, j int, lines []string) bool {
	if j+3 >= len(lines[0]) {
		return false
	}

	if lines[i][j] == 'X' {
		return lines[i][j+1] == 'M' && lines[i][j+2] == 'A' && lines[i][j+3] == 'S'
	}

	if lines[i][j] == 'S' {
		return lines[i][j+1] == 'A' && lines[i][j+2] == 'M' && lines[i][j+3] == 'X'
	}

	return false
}

func vertical(i, j int, lines []string) bool {
	if i+3 >= len(lines) {
		return false
	}

	if lines[i][j] == 'X' {
		return lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S'
	}

	if lines[i][j] == 'S' {
		return lines[i+1][j] == 'A' && lines[i+2][j] == 'M' && lines[i+3][j] == 'X'
	}

	return false
}

func diagonal_right(i, j int, lines []string) bool {
	if i+3 >= len(lines) || j+3 >= len(lines[0]) {
		return false
	}

	if lines[i][j] == 'X' {
		return lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S'
	}

	if lines[i][j] == 'S' {
		return lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'M' && lines[i+3][j+3] == 'X'
	}

	return false
}

func diagonal_left(i, j int, lines []string) bool {
	if i+3 >= len(lines) || j-3 < 0 {
		return false
	}

	if lines[i][j] == 'X' {
		return lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S'
	}

	if lines[i][j] == 'S' {
		return lines[i+1][j-1] == 'A' && lines[i+2][j-2] == 'M' && lines[i+3][j-3] == 'X'
	}

	return false
}

func cross_mas(i, j int, lines []string) bool {
	if i+2 >= len(lines) || j+2 >= len(lines[0]) {
		return false
	}

	if lines[i+1][j+1] != 'A' {
		return false
	}

	if !(lines[i][j] == 'M' && lines[i+2][j+2] == 'S' || lines[i][j] == 'S' && lines[i+2][j+2] == 'M') {
		return false
	}

	if !(lines[i][j+2] == 'M' && lines[i+2][j] == 'S' || lines[i][j+2] == 'S' && lines[i+2][j] == 'M') {
		return false
	}

	return true
}
