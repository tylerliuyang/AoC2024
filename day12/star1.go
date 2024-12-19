package main

import (
	"bufio"
)

func discoverNeighbors(visited map[[2]int]bool, i int, j int) int {
	diff := 0
	if visited[[2]int{i + 1, j}] {
		diff += 2
	}
	if visited[[2]int{i - 1, j}] {
		diff += 2
	}
	if visited[[2]int{i, j + 1}] {
		diff += 2
	}
	if visited[[2]int{i, j - 1}] {
		diff += 2
	}
	return diff
}

func bfs(area int, perimeter int, visited map[[2]int]bool, i int, j int, m [][]rune, globalvisited map[[2]int]bool) [2]int {
	if visited[[2]int{i, j}] {
		return [2]int{area, perimeter}
	}
	visited[[2]int{i, j}] = true
	globalvisited[[2]int{i, j}] = true
	perimeter += 4 - discoverNeighbors(visited, i, j)
	area += 1
	if i+1 < len(m) && m[i][j] == m[i+1][j] {
		out := bfs(area, perimeter, visited, i+1, j, m, globalvisited)
		area = out[0]
		perimeter = out[1]
	}
	if i-1 >= 0 && m[i][j] == m[i-1][j] {
		out := bfs(area, perimeter, visited, i-1, j, m, globalvisited)
		area = out[0]
		perimeter = out[1]
	}
	if j+1 < len(m[0]) && m[i][j] == m[i][j+1] {
		out := bfs(area, perimeter, visited, i, j+1, m, globalvisited)
		area = out[0]
		perimeter = out[1]
	}
	if j-1 >= 0 && m[i][j] == m[i][j-1] {
		out := bfs(area, perimeter, visited, i, j-1, m, globalvisited)
		area = out[0]
		perimeter = out[1]
	}

	return [2]int{area, perimeter}
}

func star1(scanner *bufio.Scanner) int {
	arr := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		innerarr := []rune{}
		for _, a := range line {
			innerarr = append(innerarr, a)
		}
		arr = append(arr, innerarr)
	}

	globalvisited := make(map[[2]int]bool)

	i, j := 0, 0
	total := 0
	for i < len(arr) && j < len(arr[0]) {
		if !globalvisited[[2]int{i, j}] {
			visited := make(map[[2]int]bool)
			out := bfs(0, 0, visited, i, j, arr, globalvisited)
			area := out[0]
			perimeter := out[1]
			total += area * perimeter
		}
		i += 1
		if i == len(arr) {
			i = 0
			j += 1
		}
	}

	print(total)

	return total
}
