package main

import (
	"bufio"
	"fmt"
)

func isvalidMASMAS(arr [][]rune, i int, j int) int {
	result := 0
	if (i-1) < 0 || (j-1) < 0 || (i+1) >= len(arr) || (j+1) >= len(arr[0]) {
		return 0
	}
	if arr[i][j] != 'A' {
		return 0
	}
	if arr[i-1][j-1] == 'M' && arr[i+1][j+1] == 'S' {
		result++
	}
	if arr[i-1][j-1] == 'S' && arr[i+1][j+1] == 'M' {
		result++
	}
	if arr[i-1][j+1] == 'M' && arr[i+1][j-1] == 'S' {
		result++
	}
	if arr[i-1][j+1] == 'S' && arr[i+1][j-1] == 'M' {
		result++
	}
	if arr[i+1][j-1] == 'M' && arr[i-1][j+1] == 'S' {
		result++
	}
	if arr[i+1][j-1] == 'S' && arr[i-1][j+1] == 'M' {
		result++
	}
	if arr[i+1][j+1] == 'M' && arr[i-1][j-1] == 'S' {
		result++
	}
	if arr[i+1][j+1] == 'S' && arr[i-1][j-1] == 'M' {
		result++
	}
	if result > 2 {
		return 1
	}
	return 0
}

func star2(scanner *bufio.Scanner) int {
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
			total += isvalidMASMAS(arr, i, j)
		}
	}

	fmt.Println(total)
	return total
}
