package solutions

import (
	"fmt"
	"os"
	"strings"
)

func Test12() {
	fmt.Println("Day 12 Test")
	fmt.Println("Part 1:", d12part1("./data/test12.txt"), "| Expected: 1930")
	fmt.Println("Part 2:", d12part2("./data/test12.txt"), "| Expected: 1206")
	fmt.Println()
}

func Day12() {
	fmt.Println("Day 12")
	fmt.Println("Part 1:", d12part1("./data/day12.txt"))
	fmt.Println("Part 2:", d12part2("./data/day12.txt"))
	fmt.Println()
}

func d12parseInput(filepath string) [][]rune {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	trimmed := strings.TrimSpace(string(fileContent))
	lines := strings.Split(trimmed, "\n")

	farm := [][]rune{}
	for _, line := range lines {
		farm_row := []rune{}
		for _, char := range line {
			farm_row = append(farm_row, char)
		}
		farm = append(farm, farm_row)
	}

	return farm
}

func d12part1(filename string) int {
	farm := d12parseInput(filename)
	regions := create_all_regions(farm)

	count := 0
	for _, region := range regions {
		area := region_area(region)
		perimeter := region_perimeter(region)
		count += area * perimeter
	}
	return count
}

func d12part2(filename string) int {
	farm := d12parseInput(filename)
	regions := create_all_regions(farm)

	count := 0
	for _, region := range regions {
		area := region_area(region)
		edges := horizontal_edges(region) + vertical_edges(region)
		count += area * edges
	}
	return count
}

func create_region(farm [][]rune) [][]bool {
	region := make([][]bool, len(farm))
	for i, line := range farm {
		region[i] = make([]bool, len(line))
	}
	return region
}

func get_region(farm [][]rune, region [][]bool, plant rune, x, y int) {
	if region[y][x] {
		return
	}

	if farm[y][x] == plant {
		region[y][x] = true

		bound_y := len(farm)
		bound_x := len(farm[0])
		if x > 0 {
			get_region(farm, region, plant, x-1, y)
		}
		if x < bound_x-1 {
			get_region(farm, region, plant, x+1, y)
		}
		if y > 0 {
			get_region(farm, region, plant, x, y-1)
		}
		if y < bound_y-1 {
			get_region(farm, region, plant, x, y+1)
		}
	}
}

func create_all_regions(farm [][]rune) [][][]bool {
	regions := [][][]bool{}

	for y, line := range farm {
	loop_skip:
		for x, plant := range line {
			for _, region := range regions {
				if region[y][x] {
					continue loop_skip
				}
			}
			new_region := create_region(farm)
			get_region(farm, new_region, plant, x, y)
			regions = append(regions, new_region)
		}
	}

	return regions
}

func region_area(region [][]bool) int {
	count := 0
	for _, line := range region {
		for _, b := range line {
			if b {
				count++
			}
		}
	}
	return count
}

func region_perimeter(region [][]bool) int {
	count := 0
	bound_y := len(region)
	bound_x := len(region[0])
	for y, line := range region {
		for x, b := range line {
			if !b {
				continue
			}
			if x == 0 || region[y][x-1] == false {
				count++
			}
			if x == bound_x-1 || region[y][x+1] == false {
				count++
			}
			if y == 0 || region[y-1][x] == false {
				count++
			}
			if y == bound_y-1 || region[y+1][x] == false {
				count++
			}
		}
	}
	return count
}

func horizontal_edges(region [][]bool) int {
	edges := make([][]int, len(region)+1)
	for y := 0; y < len(region)+1; y++ {
		edges[y] = make([]int, len(region[0]))
	}

	for y, line := range region {
		for x, b := range line {
			if !b {
				continue
			}

			if y == 0 || !region[y-1][x] {
				edges[y][x] = 1
			}
			if y == len(line)-1 || !region[y+1][x] {
				edges[y+1][x] = 2
			}
		}
	}

	count := 0
	for _, row := range edges {
		last_edge := 0
		for _, edge := range row {
			if edge != last_edge && edge != 0 {
				count++
			}

			last_edge = edge
		}
	}
	return count
}

func vertical_edges(region [][]bool) int {
	edges := make([][]int, len(region))
	for i, row := range region {
		edges[i] = make([]int, len(row)+1)
	}

	for x := 0; x < len(region[0]); x++ {
		for y := 0; y < len(region); y++ {
			if !region[y][x] {
				continue
			}

			if x == 0 || !region[y][x-1] {
				edges[y][x] = 1
			}
			if x == len(region[0])-1 || !region[y][x+1] {
				edges[y][x+1] = 2
			}
		}
	}

	count := 0
	for x := 0; x < len(region[0])+1; x++ {
		last_edge := 0
		for y := 0; y < len(region); y++ {
			edge := edges[y][x]
			if edge != last_edge && edge != 0 {
				count++
			}

			last_edge = edge
		}
	}
	return count
}
