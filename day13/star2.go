package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func star2(scanner *bufio.Scanner) int {
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

	final := float64(0)

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
			prize = Prize{10000000000000 + x, 10000000000000 + y}

			// Create the matrix A and vector b.
			// In this example, we will solve the equations
			// 2x + y = 3
			// x - 3y = 5
			A := mat.NewDense(2, 2, []float64{float64(buttonA.X), float64(buttonB.X), float64(buttonA.Y), float64(buttonB.Y)})
			b := mat.NewVecDense(2, []float64{float64(prize.X), float64(prize.Y)})

			// Solve the equations using the Solve function.
			var x mat.VecDense
			if err := x.SolveVec(A, b); err != nil {
				fmt.Println(err)
				continue
			}

			if math.Abs(x.AtVec(0)-float64(math.Round(x.AtVec(0)))) < 0.001 && math.Abs(x.AtVec(1)-float64(math.Round(x.AtVec(1)))) < 0.001 {
				fmt.Printf("Solution: x1 = %d, x2 = %d\n", int(x.AtVec(0)), int(x.AtVec(1)))
				final += math.Round(x.AtVec(0))*3 + math.Round(x.AtVec(1))*1
			}

			// Print the solution vector x.
			fmt.Println(x) // Output: [1 2]
		}
	}

	fmt.Println(int(final))

	return int(final)

}
