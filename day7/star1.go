package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func recurse(total int, list []int, target int) bool {
	if len(list) == 0 {
		return total == target
	}

	first, rest := list[0], list[1:]
	return recurse(total+first, rest, target) || recurse(total*first, rest, target)
}

func star1(scanner *bufio.Scanner) int {
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		var num1 int
		var str string

		parts := strings.Split(line, ":")
		str = parts[1][1:]
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			break
		}

		nums := strings.Split(str, " ")

		list := []int{}
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error parsing line:", err)
				continue
			}
			list = append(list, n)
		}
		if recurse(0, list, num1) {
			total += num1
		}
	}

	print(total)

	return total
}
