package main

import (
	"bufio"
	"fmt"
	"strings"
)

func star1(scanner *bufio.Scanner) int {
	type Button struct {
		X int
		Y int
	}

	type Prize struct {
		X int
		Y int
	}

	var buttonA Button
	var buttonB Button
	var prize Prize

	final := 0

	for scanner.Scan() {
		line := scanner.Text()
		var x, y int
		if strings.HasPrefix(line, "Button A") {
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &x, &y)
			buttonA = Button{x, y}
		} else if strings.HasPrefix(line, "Button B") {
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &x, &y)
			buttonB = Button{x, y}
		} else if strings.HasPrefix(line, "Prize") {
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &x, &y)
			prize = Prize{x, y}

			fmt.Println(buttonA, buttonB, prize)
			minPrice := 100000000
			for i := 0; i < 200; i++ {
				for j := 0; j <= i; j++ {
					newX := buttonA.X*j + buttonB.X*(i-j)
					newY := buttonA.Y*j + buttonB.Y*(i-j)
					if newX == prize.X && newY == prize.Y {
						minPrice = min(minPrice, j*3+(i-j)*1)
					}
				}
			}
			if minPrice == 100000000 {
				minPrice = 0
			}
			final += minPrice
		}
	}

	fmt.Println(final)

	return final

}
