package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <star>")
		os.Exit(1)
	}

	star := os.Args[1]

	var scanner *bufio.Scanner

	if len(os.Args) == 3 {
		file, err := os.Open(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "error opening file:", err)
			return
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	switch star {
	case "1":
		star1(scanner)
	case "2":
		star2(scanner)
	default:
		fmt.Println("Invalid star specified. Use 1 or 2.")
		os.Exit(1)
	}
}
