package main

import (
	"bufio"
	"fmt"
	"strconv"
	"unicode"
)

func star1(scanner *bufio.Scanner) int {
	total := 0
	scanner.Scan()
	line := scanner.Text()

	type StateMachine int

	const (
		E StateMachine = iota
		M
		MU
		MUL
		MULL
		MULLN
		MULLNC
		MULLNCN
	)

	SM := E

	num1 := ""
	num2 := ""
	for _, char := range line {
		switch SM {
		case E:
			SM = E
			if char == 'm' {
				num1 = ""
				num2 = ""
				SM = M
			}
		case M:
			SM = E
			if char == 'u' {
				SM = MU
			}
		case MU:
			SM = E
			if char == 'l' {
				SM = MUL
			}
		case MUL:
			SM = E
			if char == '(' {
				SM = MULL
			}
		case MULL:
			SM = E
			if unicode.IsDigit(char) {
				num1 += string(char)
				SM = MULLN
			}
		case MULLN:
			SM = E
			if unicode.IsDigit(char) {
				num1 += string(char)
				SM = MULLN
			}
			if char == ',' {
				SM = MULLNC
			}
		case MULLNC:
			SM = E
			if unicode.IsDigit(char) {
				num2 += string(char)
				SM = MULLNCN
			}
		case MULLNCN:
			SM = E
			if unicode.IsDigit(char) {
				num2 += string(char)
				SM = MULLNCN
			}
			if char == ')' {
				v1, err := strconv.Atoi(num1)
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					continue
				}
				v2, err := strconv.Atoi(num2)
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					continue
				}
				total += v1 * v2
			}
		}
	}

	fmt.Println(total)
	return total
}
