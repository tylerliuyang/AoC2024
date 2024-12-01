package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/adam-lavrik/go-imath/i64" // Functions for i64 type
)

func star1(scanner *bufio.Scanner) int64 {
	list1 := []int{}
	list2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		var num1, num2 int

		_, err := fmt.Sscanf(line, "%d %d", &num1, &num2)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error parsing line:", err)
			continue
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	total := int64(0)

	for i := 0; i < len(list1); i++ {
		total += i64.Abs(int64(list1[i]) - int64(list2[i]))
	}

	fmt.Println(total)
	return total
}
