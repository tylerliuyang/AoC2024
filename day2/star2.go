package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	// Functions for i64 type
)

func isSafe(nums []int) bool {
	var sign int
	if nums[1] > nums[0] {
		sign = -1
	} else if nums[0] > nums[1] {
		sign = 1
	} else {
		return false
	}
	good := true
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i] - nums[i+1]
		if (sign * diff) < (1) {
			good = false
			break
		} else if (sign * diff) > (3) {
			good = false
			break
		}
	}
	return good
}

func star2(scanner *bufio.Scanner) int {
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

		for i := 0; i < len(list1); i++ {
			list2 := append([]int(nil), list1[:i]...)
			list2 = append(list2, list1[i+1:]...)
			if isSafe(list2) {
				safe++
				break
			}
		}
	}

	fmt.Println(safe)
	return safe
}
