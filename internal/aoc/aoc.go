// Package aoc holds helpers shared by all the daily solutions.
package aoc

import (
	"os"
	"strings"
)

// ReadFile returns the whole contents of the file at path.
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// ReadLines returns the file's lines, without trailing newlines
// (and without a final empty line).
func ReadLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	text := strings.TrimRight(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	if text == "" {
		return nil, nil
	}
	return strings.Split(text, "\n"), nil
}
