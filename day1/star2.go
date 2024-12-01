package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	// Functions for i64 type
)

func star2(scanner *bufio.Scanner) int64 {
	list1 := []int{}
	hashMap := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()

		var num1, num2 int

		_, err := fmt.Sscanf(line, "%d %d", &num1, &num2)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error parsing line:", err)
			continue
		}
		list1 = append(list1, num1)

		hashMap[num2] += 1
	}
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	total := 0

	for i := 0; i < len(list1); i++ {
		total += list1[i] * hashMap[list1[i]]
	}

	fmt.Println(total)
	return int64(total)
}
