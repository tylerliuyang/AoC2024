package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func star1(scanner *bufio.Scanner) int {
	total := 0
	hashMap := make(map[int]map[int]struct{})

	for scanner.Scan() {
		line := scanner.Text()

		var num1, num2 int

		_, err := fmt.Sscanf(line, "%d|%d", &num1, &num2)
		if err != nil {
			break
		}
		if _, ok := hashMap[num2]; !ok {
			hashMap[num2] = make(map[int]struct{})
		}
		hashMap[num2][num1] = struct{}{}
	}

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		list1 := []int{}
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error parsing line:", err)
				continue
			}
			list1 = append(list1, n)
		}

		success := 1
		for i := 0; i < len(list1); i++ {
			for j := i + 1; j < len(list1); j++ {
				if _, ok := hashMap[list1[i]][list1[j]]; ok {
					success = 0
					break
				}
			}
		}
		total += success * list1[len(list1)/2]

	}

	print(total)

	return total
}
