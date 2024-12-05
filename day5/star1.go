package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func topoSort(outNeighbors map[int][]int) map[int]int {
	indegree := make(map[int]int)
	for _, neighbors := range outNeighbors {
		for _, neighbor := range neighbors {
			indegree[neighbor]++
		}
	}

	stack := []int{}
	for node := range outNeighbors {
		if indegree[node] == 0 {
			stack = append(stack, node)
		}
	}

	order := make(map[int]int)
	index := 0

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order[node] = index
		index++

		for _, neighbor := range outNeighbors[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				stack = append(stack, neighbor)
			}
		}
	}

	fmt.Println(order)

	return order
}

func star1(scanner *bufio.Scanner) int {
	total := 0
	hashMap := make(map[int][]int)
	allValues := make(map[int]struct{})

	for scanner.Scan() {
		line := scanner.Text()

		var num1, num2 int

		_, err := fmt.Sscanf(line, "%d|%d", &num1, &num2)
		if err != nil {
			break
		}
		hashMap[num1] = append(hashMap[num1], num2)
		allValues[num1] = struct{}{}
	}
	for v := range allValues {
		_, exists := hashMap[v]
		if !exists {
			hashMap[v] = []int{}
		}
	}

	result := topoSort(hashMap)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		list1 := []int{}
		list2 := []int{}
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error parsing line:", err)
				continue
			}
			list1 = append(list1, n)
			list2 = append(list2, result[n])
		}
		fmt.Println("list1:", list1)
		fmt.Println("list2:", list2)
		success := 1
		for i := 0; i < len(list2)-1; i++ {
			if list2[i] >= list2[i+1] {
				success = 0
			}
		}
		total += success * list1[len(list1)/2]

	}

	print(total)

	return total
}
