package main

import (
	"bufio"
	"fmt"
	"os"
)

func star2(scanner *bufio.Scanner) int {
	points := []struct {
		pX, pY, vX, vY int
	}{}

	width := 101
	height := 103
	for scanner.Scan() {
		var pX, pY, vX, vY int
		line := scanner.Text()
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &pX, &pY, &vX, &vY)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			continue
		}
		points = append(points, struct {
			pX, pY, vX, vY int
		}{pX, pY, vX, vY})
	}

	for i := 0; true; i++ {
		for j := 0; j < len(points); j++ {
			points[j].pX = (points[j].vX + points[j].pX + width) % width
			points[j].pY = (points[j].vY + points[j].pY + height) % height
		}

		grid := make([][]bool, height)
		for i := range grid {
			grid[i] = make([]bool, width)
		}

		for _, point := range points {
			grid[point.pY][point.pX] = true
		}

		found := false
		for y := 0; y < height; y++ {
			count := 0
			for x := 0; x < width; x++ {
				if grid[y][x] {
					count++
					if count == 7 {
						found = true
					}
				} else {
					count = 0
				}
			}
		}

		if found {
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					if grid[y][x] {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println(i + 1)
			}

			fmt.Print("Press 'Enter' to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}

	}

	return 0
}
