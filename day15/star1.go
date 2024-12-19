package main

import (
	"bufio"
)

func move(board [][]rune, dx int, dy int, loc struct{ x, y int }) bool {
	if loc.y+dy < 0 || loc.y+dy >= len(board[loc.x]) || loc.x+dx < 0 || loc.x+dx >= len(board) {
		return false
	}
	if board[loc.x+dx][loc.y+dy] == '#' {
		return false
	}
	if board[loc.x+dx][loc.y+dy] == '.' {
		board[loc.x+dx][loc.y+dy] = board[loc.x][loc.y]
		return true
	}
	if move(board, dx, dy, struct{ x, y int }{loc.x + dx, loc.y + dy}) {
		board[loc.x+dx][loc.y+dy] = board[loc.x][loc.y]
		return true
	}

	return false
}

func star1(scanner *bufio.Scanner) int {
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
			}
			row = append(row, char)
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
			dx = -1
		case '>':
			dy = 1
		case '<':
			dy = -1
		case 'v':
			dx = 1
		}
		if move(board, dx, dy, loc) {
			board[loc.x][loc.y] = '.'
			loc.x += dx
			loc.y += dy
		}
	}

	total := 0
	for i, row := range board {
		for j, char := range row {
			if char == 'O' {
				total += (i)*100 + (j)
			}
		}
	}

	print(total)

	return total
}
