package main

import (
	"fmt"

	"AdventOfCode2026/internal/aoc"
)

func main() {
	lines, err := aoc.ReadLines("test_input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(len(lines))
}
