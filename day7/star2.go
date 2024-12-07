package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func recurse2(total int, list []int, target int) bool {
	if len(list) == 0 {
		return total == target
	}

	first, rest := list[0], list[1:]
	potentialConcat := strconv.Itoa(total) + strconv.Itoa(first)
	concat, err := strconv.Atoi(potentialConcat)
	if err != nil {
		panic(err)
	}
	return recurse2(total+first, rest, target) || recurse2(total*first, rest, target) || recurse2(concat, rest, target)
}

func star2(scanner *bufio.Scanner) int {
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
		if recurse2(0, list, num1) {
			total += num1
		}
	}

	print(total)

	return total
}
