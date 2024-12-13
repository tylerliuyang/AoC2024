package main

import (
	"bufio"
	"strconv"
)

func bfs(graph [][]int, i int, j int, curr int) [][2]int {
	if curr == 9 {
		return [][2]int{{i, j}}
	}
	reachable := [][2]int{}
	if i+1 < len(graph) && graph[i+1][j] == (curr+1) {
		reachable = append(reachable, bfs(graph, i+1, j, curr+1)...)
	}
	if i-1 >= 0 && graph[i-1][j] == (curr+1) {
		reachable = append(reachable, bfs(graph, i-1, j, curr+1)...)
	}
	if j+1 < len(graph[0]) && graph[i][j+1] == (curr+1) {
		reachable = append(reachable, bfs(graph, i, j+1, curr+1)...)
	}
	if j-1 >= 0 && graph[i][j-1] == (curr+1) {
		reachable = append(reachable, bfs(graph, i, j-1, curr+1)...)
	}
	return reachable
}

func star1(scanner *bufio.Scanner) int {
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
		uniqueReachable := make(map[[2]int]bool)
		for _, pos := range reachable {
			uniqueReachable[pos] = true
		}
		reachable = [][2]int{}
		for pos := range uniqueReachable {
			reachable = append(reachable, pos)
		}
		total += len(reachable)
	}

	print(total)

	return total
}
