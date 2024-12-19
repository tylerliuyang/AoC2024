package main

import (
	"bufio"
	"fmt"
)

func star1(scanner *bufio.Scanner) int {
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

	for i := 0; i < 100; i++ {
		for j := 0; j < len(points); j++ {
			points[j].pX = (points[j].vX + points[j].pX + width) % width
			points[j].pY = (points[j].vY + points[j].pY + height) % height
		}
		fmt.Println(points[0])
	}

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	for _, point := range points {
		if point.pX < width/2 && point.pY < height/2 {
			q1++
		} else if point.pX > width/2 && point.pY < height/2 {
			q2++
		} else if point.pX < width/2 && point.pY > height/2 {
			q3++
		} else if point.pX > width/2 && point.pY > height/2 {
			q4++
		} else {
			fmt.Println(point, "erorr")
		}
	}

	print(q1 * q2 * q3 * q4)

	return q1 * q2 * q3 * q4

}
