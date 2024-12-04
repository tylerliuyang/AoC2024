package main

import (
	"bufio"
	"fmt"
)

func isValidXMAS(arr [][]rune, i int, j int) int {
	result := 0
	if (i - 3) >= 0 {
		if arr[i-3][j] == 'S' && arr[i-2][j] == 'A' && arr[i-1][j] == 'M' && arr[i][j] == 'X' {
			result++
		}
	}
	if (i + 3) < len(arr) {
		if arr[i][j] == 'X' && arr[i+1][j] == 'M' && arr[i+2][j] == 'A' && arr[i+3][j] == 'S' {
			result++
		}
	}
	if (j - 3) >= 0 {
		if arr[i][j-3] == 'S' && arr[i][j-2] == 'A' && arr[i][j-1] == 'M' && arr[i][j] == 'X' {
			result++
		}
	}
	if (j + 3) < len(arr[0]) {
		if arr[i][j] == 'X' && arr[i][j+1] == 'M' && arr[i][j+2] == 'A' && arr[i][j+3] == 'S' {
			result++
		}
	}
	if (i-3) >= 0 && (j-3) >= 0 {
		if arr[i-3][j-3] == 'S' && arr[i-2][j-2] == 'A' && arr[i-1][j-1] == 'M' && arr[i][j] == 'X' {
			result++
		}
	}
	if (i+3) < len(arr) && (j+3) < len(arr[0]) {
		if arr[i][j] == 'X' && arr[i+1][j+1] == 'M' && arr[i+2][j+2] == 'A' && arr[i+3][j+3] == 'S' {
			result++
		}
	}
	if (i-3) >= 0 && (j+3) < len(arr[0]) {
		if arr[i-3][j+3] == 'S' && arr[i-2][j+2] == 'A' && arr[i-1][j+1] == 'M' && arr[i][j] == 'X' {
			result++
		}
	}
	if (i+3) < len(arr) && (j-3) >= 0 {
		if arr[i][j] == 'X' && arr[i+1][j-1] == 'M' && arr[i+2][j-2] == 'A' && arr[i+3][j-3] == 'S' {
			result++
		}
	}
	return result
}

func star1(scanner *bufio.Scanner) int {
	total := 0
	arr := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		innerarr := make([]rune, 0)
		for _, a := range line {
			innerarr = append(innerarr, a)
		}
		arr = append(arr, innerarr)
	}
	for _, row := range arr {
		for _, char := range row {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			total += isValidXMAS(arr, i, j)
		}
	}

	fmt.Println(total)
	return total
}
