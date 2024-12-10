package main

import (
	"bufio"
	"fmt"
)

func findLengthOfRegion(i int, arr []int) int {
	count := 0
	for _, val := range arr {
		if val == i {
			count++
		}
	}
	return count
}

func findFirstIndexOfSizeRegion(desiredSize int, arr []int) int {
	size := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == -1 {
			size += 1
		} else {
			size = 0
		}
		if size == desiredSize {
			return i - size + 1
		}
	}
	return -1
}

func findNumber(arr []int, i int) int {
	index := 0
	for _, val := range arr {
		if val == i {
			return index
		}
		index += 1
	}
	return -1
}

func star2(scanner *bufio.Scanner) int {
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

	start := i/2 + 1
	print(start)

	for start > 0 {
		start--
		length := findLengthOfRegion(start, sparseRepresentation)
		firstIndex := findFirstIndexOfSizeRegion(length, sparseRepresentation)
		index := findNumber(sparseRepresentation, start)
		if index < firstIndex {
			continue
		}
		if firstIndex == -1 {
			continue
		}

		for k := firstIndex; k < firstIndex+length; k++ {
			sparseRepresentation[k] = start
		}

		for k := index; k < index+length; k++ {
			sparseRepresentation[k] = -1
		}
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
