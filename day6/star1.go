package main

import (
	"bufio"
)

func star1(scanner *bufio.Scanner) int {
	total := 0
	location := []int{0, 0}
	hashMap := [][]int{}
	visited := make(map[[2]int]bool)
	currentDirection := [2]int{0, 0}

	for scanner.Scan() {
		line := scanner.Text()
		hashMap = append(hashMap, []int{})
		for _, num := range line {
			hashMap[len(hashMap)-1] = append(hashMap[len(hashMap)-1], int(num))
			if num == '^' || num == 'v' || num == '<' || num == '>' {
				location = []int{len(hashMap) - 1, len(hashMap[len(hashMap)-1]) - 1}
				if num == '^' {
					currentDirection = [2]int{-1, 0}
				} else if num == 'v' {
					currentDirection = [2]int{1, 0}
				} else if num == '<' {
					currentDirection = [2]int{0, -1}
				} else if num == '>' {
					currentDirection = [2]int{0, 1}
				}
			}
		}
	}

	rows := len(hashMap)
	cols := len(hashMap[0])

	for location != nil {
		if _, ok := visited[[2]int{location[0], location[1]}]; !ok {
			total++
			visited[[2]int{location[0], location[1]}] = true
		}
		if location[0]+currentDirection[0] < 0 || location[0]+currentDirection[0] >= rows || location[1]+currentDirection[1] < 0 || location[1]+currentDirection[1] >= cols {
			break
		}
		collided := hashMap[location[0]+currentDirection[0]][location[1]+currentDirection[1]] == '#'
		if collided {
			if currentDirection == [2]int{-1, 0} {
				currentDirection = [2]int{0, 1}
			} else if currentDirection == [2]int{1, 0} {
				currentDirection = [2]int{0, -1}
			} else if currentDirection == [2]int{0, -1} {
				currentDirection = [2]int{-1, 0}
			} else if currentDirection == [2]int{0, 1} {
				currentDirection = [2]int{1, 0}
			}
			continue
		} else {
			location = []int{location[0] + currentDirection[0], location[1] + currentDirection[1]}
		}
	}

	print(total)

	return total
}
