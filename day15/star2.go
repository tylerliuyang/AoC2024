package main

import (
	"bufio"
)

func move2(board [][]rune, dx int, dy int, loc struct{ x, y int }) bool {
	if loc.y+dy < 0 || loc.y+dy >= len(board) || loc.x+dx < 0 || loc.x+dx >= len(board[loc.y]) {
		return false
	}
	dest := board[loc.y+dy][loc.x+dx]
	source := board[loc.y][loc.x]
	if dest == '#' {
		return false
	}
	if dest == '.' {
		return true
	}
	if dest == '[' {
		if dy != 0 {
			leftBox := move2(board, dx, dy, struct{ x, y int }{loc.x, loc.y + dy})
			rightBox := move2(board, dx, dy, struct{ x, y int }{loc.x + 1, loc.y + dy})
			if leftBox && rightBox {
				board[loc.y+dy+dy][loc.x] = '['
				board[loc.y+dy+dy][loc.x+1] = ']'
				board[loc.y+dy][loc.x] = '.'
				board[loc.y+dy][loc.x+1] = '.'
				return true
			}
		} else {
			if move2(board, dx, dy, struct{ x, y int }{loc.x + 2, loc.y}) {
				board[loc.y][loc.x+3] = ']'
				board[loc.y][loc.x+2] = '['
				board[loc.y][loc.x+1] = '.'
				return true
			}
		}

		return false
	}
	if dest == ']' {
		if dy != 0 {
			leftBox := move2(board, dx, dy, struct{ x, y int }{loc.x - 1, loc.y + dy})
			rightBox := move2(board, dx, dy, struct{ x, y int }{loc.x, loc.y + dy})
			if leftBox && rightBox {
				board[loc.y+dy+dy][loc.x-1] = '['
				board[loc.y+dy+dy][loc.x] = ']'
				board[loc.y+dy][loc.x-1] = '.'
				board[loc.y+dy][loc.x] = '.'
				return true
			}
		} else {
			if move2(board, dx, dy, struct{ x, y int }{loc.x - 2, loc.y + dy}) {
				board[loc.y][loc.x-3] = '['
				board[loc.y][loc.x-2] = ']'
				board[loc.y][loc.x-1] = '.'
				return true
			}
			return false
		}
	}
	if dest == '@' {
		if move2(board, dx, dy, struct{ x, y int }{loc.x + dx, loc.y + dy}) {
			board[loc.y+dy][loc.x+dx] = source
			board[loc.y][loc.x] = '.'
			return true
		}
		return false
	}

	return false
}

func star2(scanner *bufio.Scanner) int {
	var board [][]rune
	loc := struct {
		x, y int
	}{}

	for scanner.Scan() {
		var row []rune
		line := scanner.Text()
		if line == "" {
			break
		}
		for _, char := range line {
			if char == '@' {
				loc.x = len(row)
				loc.y = len(board)
				row = append(row, '@')
				row = append(row, '.')
			}
			if char == 'O' {
				row = append(row, '[')
				row = append(row, ']')
			}
			if char == '#' {
				row = append(row, '#')
				row = append(row, '#')
			}
			if char == '.' {
				row = append(row, '.')
				row = append(row, '.')
			}
		}
		board = append(board, row)
	}

	var commands []rune
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			commands = append(commands, char)
		}
	}

	for _, command := range commands {
		dx := 0
		dy := 0
		switch command {
		case '^':
			dy = -1
		case '>':
			dx = 1
		case '<':
			dx = -1
		case 'v':
			dy = 1
		}
		copyBoard := make([][]rune, len(board))
		for i, row := range board {
			copyBoard[i] = make([]rune, len(row))
			copy(copyBoard[i], row)
		}

		if move2(board, dx, dy, loc) {
			board[loc.y+dy][loc.x+dx] = '@'
			board[loc.y][loc.x] = '.'
			loc.x += dx
			loc.y += dy
		} else {
			board = copyBoard
		}

	}

	total := 0
	for i, row := range board {
		for j, char := range row {
			if char == '[' {
				total += (i)*100 + (j)
			}
		}
	}

	print(total)

	return total
}
