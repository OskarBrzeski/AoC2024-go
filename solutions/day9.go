package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test9() {
	fmt.Println("Day 9 Test")
	fmt.Println("Part 1:", d9part1("./data/test9.txt"), "| Expected: 1928")
	fmt.Println("Part 2:", d9part2("./data/test9.txt"), "| Expected: 2858")
	fmt.Println()
}

func Day9() {
	fmt.Println("Day 9")
	fmt.Println("Part 1:", d9part1("./data/day9.txt"))
	fmt.Println("Part 2:", d9part2("./data/day9.txt"))
	fmt.Println()
}

func d9parseInput(filepath string) []int {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	trimmed := strings.TrimSpace(string(fileContent))

	nums := make([]int, len(trimmed))
	for i, char := range trimmed {
		num, _ := strconv.Atoi(string(char))
		nums[i] = num
	}

	return nums
}

func d9part1(filename string) int {
	nums := d9parseInput(filename)
	disk := create_disk(nums)
	compact := compact_disk(disk)
	count := 0
	for i, num := range compact {
		count += i * num
	}
	return count
}

type file struct {
	pos  int
	size int
	id   int
}

func d9part2(filename string) int {
	nums := d9parseInput(filename)
	files := create_files(nums)
	compact := compact_files(files)
	disk := files_to_disk(compact)

	count := 0
	for i, num := range disk {
		if num != -1 {
			count += i * num
		}
	}
	return count
}

func create_disk(nums []int) []int {
	length := 0
	for _, num := range nums {
		length += num
	}

	disk := make([]int, length)
	index := 0
	for i, num := range nums {
		for j := 0; j < num; j++ {
			if i%2 == 1 {
				disk[index] = -1
			} else {
				disk[index] = i / 2
			}
			index++
		}
	}

	return disk
}

func compact_disk(disk []int) []int {
	left := 0
	right := len(disk) - 1

	for left < right {
		if disk[right] == -1 {
			right--
			continue
		}
		if disk[left] == -1 {
			disk[left] = disk[right]
			right--
		}
		left++
	}

	return disk[:right+1]
}

func create_files(nums []int) []file {
	files := make([]file, (len(nums)+1)/2)

	pos := 0
	for i, num := range nums {
		if i%2 == 0 {
			files[i/2] = file{pos, num, i / 2}
		}

		pos += num
	}

	return files
}

func compact_files(old_files []file) []file {
	files := make([]file, len(old_files))
	copy(files, old_files)

	filemove := len(files) - 1
	last_checked := len(files)

outer:
	for filemove > 0 {
		if last_checked < files[filemove].id {
			filemove--
			continue
		}
		last_checked = files[filemove].id
		for left := 1; left <= filemove; left++ {
			if files[left].pos-files[left-1].pos-files[left-1].size >= files[filemove].size {
				files[filemove].pos = files[left-1].pos + files[left-1].size

				files_new := []file{}
				files_new = append(files_new, files[:left]...)
				files_new = append(files_new, files[filemove])
				files_new = append(files_new, files[left:filemove]...)
				files_new = append(files_new, files[filemove+1:]...)
				files = files_new

				continue outer
			}
		}
		filemove--
	}

	return files
}

func files_to_disk(files []file) []int {
	size := files[len(files)-1].pos + files[len(files)-1].size
	disk := make([]int, size)

	file_index := 0
	for i := 0; i < size; i++ {
		if files[file_index].pos+files[file_index].size <= i {
			file_index++
		}

		if files[file_index].pos > i {
			disk[i] = -1
		} else {
			disk[i] = files[file_index].id
		}
	}

	return disk
}
