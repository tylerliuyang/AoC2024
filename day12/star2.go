package main

import (
	"bufio"
)

func discoverNeighbors2(visited map[[2]int]bool, i int, j int) int {
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

func bfs2(area int, perimeter int, visited map[[2]int]bool, i int, j int, m [][]rune, globalvisited map[[2]int]bool) [2]int {
	if visited[[2]int{i, j}] {
		return [2]int{area, perimeter}
	}
	visited[[2]int{i, j}] = true
	globalvisited[[2]int{i, j}] = true
	perimeter += 4 - discoverNeighbors2(visited, i, j)
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

func countSides(visited map[[2]int]bool) int {
	tiles := make(map[[3]int]rune)
	for loc, _ := range visited {
		i := loc[0]
		j := loc[1]
		if !visited[[2]int{i + 1, j}] {
			tiles[[3]int{i + 1, j, 1}] = 1
		}
		if !visited[[2]int{i, j + 1}] {
			tiles[[3]int{i, j + 1, 2}] = 1
		}
		if !visited[[2]int{i - 1, j}] {
			tiles[[3]int{i - 1, j, 3}] = 1
		}
		if !visited[[2]int{i, j - 1}] {
			tiles[[3]int{i, j - 1, 4}] = 1
		}
	}

	sides := 0

	for len(tiles) > 0 {
		queue := [][3]int{}
		for k := range tiles {
			queue = append(queue, k)
			delete(tiles, k)
			break
		}
		for len(queue) > 0 {
			i := queue[0][0]
			j := queue[0][1]
			direction := queue[0][2]
			queue = queue[1:]
			if tiles[[3]int{i + 1, j, direction}] == 1 {
				queue = append(queue, [3]int{i + 1, j, direction})
				delete(tiles, [3]int{i + 1, j, direction})
			}
			if tiles[[3]int{i, j + 1, direction}] == 1 {
				queue = append(queue, [3]int{i, j + 1, direction})
				delete(tiles, [3]int{i, j + 1, direction})
			}
			if tiles[[3]int{i - 1, j, direction}] == 1 {
				queue = append(queue, [3]int{i - 1, j, direction})
				delete(tiles, [3]int{i - 1, j, direction})
			}
			if tiles[[3]int{i, j - 1, direction}] == 1 {
				queue = append(queue, [3]int{i, j - 1, direction})
				delete(tiles, [3]int{i, j - 1, direction})
			}
		}
		sides += 1
	}
	return sides
}

func star2(scanner *bufio.Scanner) int {
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
			out := bfs2(0, 0, visited, i, j, arr, globalvisited)
			sides := countSides(visited)
			area := out[0]
			// perimeter := out[1]
			total += area * sides
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
