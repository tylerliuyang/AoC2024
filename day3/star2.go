package main

import (
	"bufio"
	"fmt"
	"strconv"
	"unicode"
)

func star2(scanner *bufio.Scanner) int {
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

	const (
		E1 StateMachine = iota
		D
		DO
		DON
		DONQ
		DONQT
		DONQTL
		DOL
	)
	enabled := true
	TRACKSTATE := E1
	SM := E

	num1 := ""
	num2 := ""
	for _, char := range line {
		switch TRACKSTATE {
		case E1:
			TRACKSTATE = E1
			if char == 'd' {
				TRACKSTATE = D
			}
		case D:
			TRACKSTATE = E1
			if char == 'o' {
				TRACKSTATE = DO
			}
		case DO:
			TRACKSTATE = E1
			if char == 'n' {
				TRACKSTATE = DON
			}
			if char == '(' {
				TRACKSTATE = DOL
			}
		case DON:
			TRACKSTATE = E1
			if char == '\'' {
				TRACKSTATE = DONQ
			}
		case DONQ:
			TRACKSTATE = E1
			if char == 't' {
				TRACKSTATE = DONQT
			}
		case DONQT:
			TRACKSTATE = E1
			if char == '(' {
				TRACKSTATE = DONQTL
			}
		case DONQTL:
			TRACKSTATE = E1
			if char == ')' {
				TRACKSTATE = E1
				enabled = false
			}
		case DOL:
			TRACKSTATE = E1
			if char == ')' {
				TRACKSTATE = E1
				enabled = true
			}
		}
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
				if enabled {
					total += v1 * v2
				}
			}
		}
	}

	fmt.Println(total)
	return total
}
