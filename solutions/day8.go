package solutions

import (
	"fmt"
	"os"
	"strings"
)

func Test8() {
	fmt.Println("Day 8 Test")
	fmt.Println("Part 1:", d8part1("./data/test8.txt"), "| Expected: 14")
	fmt.Println("Part 2:", d8part2("./data/test8.txt"), "| Expected: 34")
	fmt.Println()
}

func Day8() {
	fmt.Println("Day 8")
	fmt.Println("Part 1:", d8part1("./data/day8.txt"))
	fmt.Println("Part 2:", d8part2("./data/day8.txt"))
	fmt.Println()
}

type location struct {
	x int
	y int
}

func d8parseInput(filepath string) (int, int, map[rune][]location) {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	trimmed := strings.TrimSpace(string(fileContent))
	lines := strings.Fields(trimmed)

	location_map := map[rune][]location{}
	for y, line := range lines {
		for x, char := range line {
			if char == '.' {
				continue
			}
			if location_map[char] == nil {
				location_map[char] = []location{}
			}
			location_map[char] = append(location_map[char], location{x, y})
		}
	}

	return len(lines), len(lines[0]), location_map
}

func d8part1(filename string) int {
	height, width, location_map := d8parseInput(filename)

	antinodes := make([]bool, width*height)
	for _, locs := range location_map {
		for i := 0; i < len(locs)-1; i++ {
			for j := i + 1; j < len(locs); j++ {
				an1, an2 := find_antinodes(locs[i], locs[j], width, height)
				if an1 != -1 {
					antinodes[an1] = true
				}
				if an2 != -1 {
					antinodes[an2] = true
				}
			}
		}
	}

	count := 0
	for _, cell := range antinodes {
		if cell {
			count++
		}
	}
	return count
}

func d8part2(filename string) int {
	height, width, location_map := d8parseInput(filename)

	antinodes := make([]bool, width*height)
	for _, locs := range location_map {
		for i := 0; i < len(locs)-1; i++ {
			for j := i + 1; j < len(locs); j++ {
				hs := find_harmonics(locs[i], locs[j], width, height)
				for _, h := range hs {
					antinodes[h] = true
				}
			}
		}
	}

	count := 0
	for _, cell := range antinodes {
		if cell {
			count++
		}
	}
	return count
}

func find_antinodes(loc1, loc2 location, width, height int) (int, int) {
	dx := loc2.x - loc1.x
	dy := loc2.y - loc1.y

	var an1, an2 int

	an1x := loc1.x - dx
	an1y := loc1.y - dy
	if an1x >= 0 && an1x < width && an1y >= 0 && an1y < height {
		an1 = an1y*width + an1x
	} else {
		an1 = -1
	}

	an2x := loc2.x + dx
	an2y := loc2.y + dy
	if an2x >= 0 && an2x < width && an2y >= 0 && an2y < height {
		an2 = an2y*width + an2x
	} else {
		an2 = -1
	}

	return an1, an2
}

func find_harmonics(loc1, loc2 location, width, height int) []int {
	result := []int{loc1.y*width + loc1.x, loc2.y*width + loc2.x}

	dx := loc2.x - loc1.x
	dy := loc2.y - loc1.y

	an1x := loc1.x - dx
	an1y := loc1.y - dy

	for an1x >= 0 && an1x < width && an1y >= 0 && an1y < height {
		result = append(result, an1y*width+an1x)
		an1x -= dx
		an1y -= dy
	}

	an2x := loc2.x + dx
	an2y := loc2.y + dy

	for an2x >= 0 && an2x < width && an2y >= 0 && an2y < height {
		result = append(result, an2y*width+an2x)
		an2x += dx
		an2y += dy
	}

	return result
}

// This function was used to print out the input with the antinodes
// and harmonics inserted so that it could be referenced with the
// example on the website.
func render_antinodes(location_map map[rune][]location, ans []bool, width, height int) {
	printed := make([]rune, width*height)
	for i := 0; i < width*height; i++ {
		if ans[i] {
			printed[i] = '#'
		} else {
			printed[i] = '.'
		}
	}

	for char, locs := range location_map {
		for _, loc := range locs {
			printed[loc.y*width+loc.x] = char
		}
	}

	for i := 0; i < height; i++ {
		fmt.Println(string(printed[i*width : (i+1)*width]))
	}
}
