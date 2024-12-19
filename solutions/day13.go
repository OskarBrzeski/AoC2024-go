package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test13() {
	fmt.Println("Day 13 Test")
	fmt.Println("Part 1:", d13part1("./data/test13.txt"), "| Expected: 480")
	fmt.Println("Part 2:", d13part2("./data/test13.txt"), "| Expected: 875318608908")
	fmt.Println()
}

func Day13() {
	fmt.Println("Day 13")
	fmt.Println("Part 1:", d13part1("./data/day13.txt"))
	fmt.Println("Part 2:", d13part2("./data/day13.txt"))
	fmt.Println()
}

type vector2 struct {
	x int
	y int
}

type clawMachine struct {
	x int
	y int
	a vector2
	b vector2
}

func d13parseInput(filepath string) []clawMachine {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	trimmed := strings.TrimSpace(string(fileContent))
	sections := strings.Split(trimmed, "\n\n")

	machines := []clawMachine{}
	for _, section := range sections {
		lines := strings.Split(section, "\n")
		ax, _ := strconv.Atoi(lines[0][12:14])
		ay, _ := strconv.Atoi(lines[0][18:20])
		bx, _ := strconv.Atoi(lines[1][12:14])
		by, _ := strconv.Atoi(lines[1][18:20])
		lastline := strings.Split(lines[2][7:], ", ")
		x, _ := strconv.Atoi(lastline[0][2:])
		y, _ := strconv.Atoi(lastline[1][2:])
		claw := clawMachine{x: x, y: y, a: vector2{x: ax, y: ay}, b: vector2{x: bx, y: by}}
		machines = append(machines, claw)
	}

	return machines
}

func d13part1(filename string) int {
	clawMachines := d13parseInput(filename)
	count := 0
	for _, claw := range clawMachines {
		count += fewestTokens(claw)
	}
	return count
}

func d13part2(filename string) int {
	clawMachine := d13parseInput(filename)
	count := 0
	for _, claw := range clawMachine {
		count += math(claw)
	}
	return count
}

func fewestTokens(claw clawMachine) int {
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			tokens := calcTokens(a, b, claw)
			if tokens != -1 {
				return tokens
			}
		}
	}

	return 0
}

func calcTokens(aCount, bCount int, claw clawMachine) int {
	xPos := aCount*claw.a.x + bCount*claw.b.x
	yPos := aCount*claw.a.y + bCount*claw.b.y
	if claw.x == xPos && claw.y == yPos {
		return aCount*3 + bCount
	}
	return -1
}

const offset = 10000000000000

func math(claw clawMachine) int {
	claw.x += offset
	claw.y += offset

	m1 := float64(claw.a.y) / float64(claw.a.x)
	m2 := float64(claw.b.y) / float64(claw.b.x)
	c := float64(claw.y) - m2*float64(claw.x)

	x := c / (m1 - m2)
	// y := m1 * x

	aPresses := x / float64(claw.a.x)
	bPresses := (float64(claw.x) - x) / float64(claw.b.x)

	for i := -100; i < 100; i++ {
		for j := -100; j < 100; j++ {
			if works(int(aPresses)+i, int(bPresses)+j, claw) {
				return (int(aPresses)+i)*3 + (int(bPresses) + j)
			}
		}
	}
	return 0
}

func works(aPresses, bPresses int, claw clawMachine) bool {
	if aPresses < 0 {
		return false
	}
	if bPresses < 0 {
		return false
	}
	if (aPresses*claw.a.x + bPresses*claw.b.x) != claw.x {
		return false
	}
	if (aPresses*claw.a.y + bPresses*claw.b.y) != claw.y {
		return false
	}
	return true
}
