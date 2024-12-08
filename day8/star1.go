package main

import (
	"bufio"
)

func star1(scanner *bufio.Scanner) int {
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
					antiNodes[[2]int{v3[0] + (v3[0] - v2[0]), v3[1] + (v3[1] - v2[1])}] = true
					antiNodes[[2]int{v2[0] - (v3[0] - v2[0]), v2[1] - (v3[1] - v2[1])}] = true
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
