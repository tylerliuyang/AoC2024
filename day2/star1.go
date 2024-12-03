package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func star1(scanner *bufio.Scanner) int {
	safe := 0

	for scanner.Scan() {
		list1 := []int{}

		line := scanner.Text()

		parts := strings.Split(line, " ")

		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			list1 = append(list1, num)
		}

		var sign int
		if list1[1] > list1[0] {
			sign = -1
		} else if list1[0] > list1[1] {
			sign = 1
		} else {
			continue
		}
		good := 1
		for i := 0; i < len(list1)-1; i++ {
			diff := list1[i] - list1[i+1]
			if (sign * diff) < (1) {
				good = 0
				break
			} else if (sign * diff) > (3) {
				good = 0
				break
			}
		}
		safe += good
	}

	fmt.Println(safe)
	return safe
}
