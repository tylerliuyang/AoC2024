package main

import (
	"bufio"
)

func rotateDirection(direction [2]int) [2]int {
	if direction == [2]int{-1, 0} {
		return [2]int{0, 1}
	} else if direction == [2]int{1, 0} {
		return [2]int{0, -1}
	} else if direction == [2]int{0, -1} {
		return [2]int{-1, 0}
	} else if direction == [2]int{0, 1} {
		return [2]int{1, 0}
	}
	return [2]int{0, 0}
}

func simulate(location []int, currentDirection [2]int, hashMap [][]int) int {
	total := 0
	rows := len(hashMap)
	cols := len(hashMap[0])
	visited := make(map[[2]int]bool)
	steps := 0

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
			currentDirection = rotateDirection(currentDirection)
			continue
		} else {
			location = []int{location[0] + currentDirection[0], location[1] + currentDirection[1]}
		}
		steps++
		if steps > rows*cols {
			return 1
		}
	}
	return 0
}

func star2(scanner *bufio.Scanner) int {
	total := 0
	hashMap := [][]int{}

	location := []int{0, 0}
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

	for i := 0; i < len(hashMap); i++ {
		for j := 0; j < len(hashMap[0]); j++ {
			if location[0] == i && location[1] == j {
				continue
			}
			temp := hashMap[i][j]
			hashMap[i][j] = '#'
			total += simulate(location, currentDirection, hashMap)
			hashMap[i][j] = temp
		}
	}

	print(total)

	return total
}
