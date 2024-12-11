package solutions

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func Test11() {
	fmt.Println("Day 11 Test")
	fmt.Println("Part 1:", d11part1("./data/test11.txt"), "| Expected: 55312")
	fmt.Println("Part 2:", d11part2("./data/test11.txt"), "| Expected: 65601038650482")
	fmt.Println()
}

func Day11() {
	fmt.Println("Day 11")
	fmt.Println("Part 1:", d11part1("./data/day11.txt"))
	fmt.Println("Part 2:", d11part2("./data/day11.txt"))
	fmt.Println()
}

func d11parseInput(filepath string) map[*big.Int]int {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	line := strings.Fields(string(fileContent))
	nums := map[*big.Int]int{}
	for _, numstr := range line {
		num, _ := new(big.Int).SetString(numstr, 10)
		nums[num] += 1
	}

	return nums
}

func d11part1(filename string) int {
	stones := d11parseInput(filename)

	for i := 0; i < 25; i++ {
		stones = blink2(stones)
	}

	count := 0

	for _, freq := range stones {
		count += freq
	}

	return count
}

func d11part2(filename string) int {
	stones := d11parseInput(filename)

	for i := 0; i < 75; i++ {
		stones = blink2(stones)
	}

	count := 0

	for _, freq := range stones {
		count += freq
	}

	return count
}

func blink(stones []*big.Int) []*big.Int {
	new_stones := make([]*big.Int, len(stones)*2)
	index := 0
	for _, stone := range stones {
		numstr := stone.Text(10)

		if stone.Cmp(big.NewInt(0)) == 0 {
			new_stones[index] = big.NewInt(1)
			index++
		} else if len(numstr)%2 == 0 {
			left, _ := new(big.Int).SetString(numstr[:len(numstr)/2], 10)
			new_stones[index] = left
			index++
			right, _ := new(big.Int).SetString(numstr[len(numstr)/2:], 10)
			new_stones[index] = right
			index++
		} else {
			new_stones[index] = new(big.Int).Mul(stone, big.NewInt(2024))
			index++
		}
	}
	return new_stones[:index]
}

func blink2(stones map[*big.Int]int) map[*big.Int]int {
	new_stones := make(map[*big.Int]int)
	for num, count := range stones {
		numstr := num.Text(10)

		if len(numstr)%2 == 0 {
			left, _ := new(big.Int).SetString(numstr[:len(numstr)/2], 10)
			index_left := find_index(new_stones, left)
			if index_left == nil {
				new_stones[left] = count
			} else {
				new_stones[index_left] += count
			}

			right, _ := new(big.Int).SetString(numstr[len(numstr)/2:], 10)
			index_right := find_index(new_stones, right)
			if index_right == nil {
				new_stones[right] = count
			} else {
				new_stones[index_right] += count
			}

		} else if num.Cmp(big.NewInt(0)) == 0 {
			one := big.NewInt(1)
			index := find_index(new_stones, one)
			if index == nil {
				new_stones[one] = count
			} else {
				new_stones[index] += count
			}

		} else {
			new_stone := new(big.Int).Mul(num, big.NewInt(2024))
			index := find_index(new_stones, new_stone)
			if index == nil {
				new_stones[new_stone] = count
			} else {
				new_stones[index] += count
			}
		}
	}

	return new_stones
}

func find_index(stones map[*big.Int]int, num *big.Int) *big.Int {
	for stone := range stones {
		if stone.Cmp(num) == 0 {
			return stone
		}
	}
	return nil
}
