package main

import (
	"bufio"
)

func star2(scanner *bufio.Scanner) int {
	total := 0
	things := make(map[rune][][2]int)
	curRow := 0
	curCol := 0

	for scanner.Scan() {
		line := scanner.Text()
		curCol = 0
		for _, num := range line {
			if num != '.' {
				things[num] = append(things[num], [2]int{curRow, curCol})
			}
			curCol++
		}
		curRow++
	}

	antiNodes := make(map[[2]int]bool)
	for _, v := range things {
		for _, v2 := range v {
			for _, v3 := range v {
				if v2 != v3 {
					movementVector := [2]int{v3[0] - v2[0], v3[1] - v2[1]}
					newAntiNode := [2]int{v2[0], v2[1]}
					newAntiNode2 := [2]int{v3[0], v3[1]}

					for (newAntiNode[0] >= 0 && newAntiNode[0] < curRow) || (newAntiNode2[1] >= 0 && newAntiNode2[1] < curCol) {
						antiNodes[newAntiNode] = true
						newAntiNode[0] -= movementVector[0]
						newAntiNode[1] -= movementVector[1]
						antiNodes[newAntiNode2] = true
						newAntiNode2[0] += movementVector[0]
						newAntiNode2[1] += movementVector[1]
					}
				}
			}
		}
	}

	for loc, _ := range antiNodes {
		if loc[0] >= 0 && loc[0] < curRow && loc[1] >= 0 && loc[1] < curCol {
			total++
		}
	}

	print(total)

	return total
}
