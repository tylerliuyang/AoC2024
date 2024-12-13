package main

import (
	"bufio"
	"strconv"
)

func star2(scanner *bufio.Scanner) int {
	starts := [][2]int{}
	arr := make([][]int, 0)
	i, j := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		innerarr := make([]int, 0)
		for _, a := range line {
			value, err := strconv.Atoi(string(a))
			if err != nil {
				panic(err)
			}
			innerarr = append(innerarr, value)
			if a == '0' {
				starts = append(starts, [2]int{i, j})
			}
			j += 1
		}
		i += 1
		j = 0
		arr = append(arr, innerarr)
	}

	total := 0

	for _, start := range starts {
		i, j := start[0], start[1]
		reachable := bfs(arr, i, j, 0)
		total += len(reachable)
	}

	print(total)

	return total
}
