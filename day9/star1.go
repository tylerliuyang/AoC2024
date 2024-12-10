package main

import (
	"bufio"
	"fmt"
)

func findLastNonNegIndex(arr []int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] >= 0 {
			return i
		}
	}
	return -1
}

func findFirstNegIndex(arr []int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			return i
		}
	}
	return -1
}

func star1(scanner *bufio.Scanner) int {
	scanner.Scan()
	line := scanner.Text()

	sparseRepresentation := []int{}
	i := 0
	for _, char := range line {
		var c int
		if i%2 == 0 {
			c = i / 2
		} else {
			c = -1
		}
		for j := 0; j < int(char-'0'); j++ {
			sparseRepresentation = append(sparseRepresentation, c)
		}
		i++
	}

	for {
		lastNonNegIndex := findLastNonNegIndex(sparseRepresentation)
		if lastNonNegIndex == -1 {
			break
		}
		firstNegIndex := findFirstNegIndex(sparseRepresentation)

		if firstNegIndex > lastNonNegIndex {
			break
		}

		sparseRepresentation[firstNegIndex] = sparseRepresentation[lastNonNegIndex]
		sparseRepresentation[lastNonNegIndex] = -1
	}

	fmt.Println(sparseRepresentation)

	total := 0
	for i := 0; i < len(sparseRepresentation); i++ {
		if sparseRepresentation[i] < 0 {
			continue
		}
		total += i * sparseRepresentation[i]
	}

	print(total)
	return 0
}
