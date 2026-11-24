package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const template = `package main

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
`

func createDay() error {
	day, name := os.Args[1], os.Args[2:]
	i, err := strconv.Atoi(day)

	if err != nil {
		return err
	}

	if len(name) == 0 {
		return errors.New("'name' is missing")
	}

	dir := fmt.Sprintf("day_%02d_%s", i, strings.ToLower(strings.Join(name, "_")))
	files := [3]string{"main.go", "test_input.txt", "input.txt"}

	for part := 1; part <= 2; part++ {
		partDir := filepath.Join(dir, fmt.Sprintf("part%d", part))
		if err := os.MkdirAll(partDir, 0o755); err != nil {
			return err
		}

		// O_EXCL: never overwrite a solution that already exists.
		for i, file := range files {
			solutionFile := filepath.Join(partDir, file)
			f, err := os.OpenFile(solutionFile, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
			if errors.Is(err, os.ErrExist) {
				continue
			}
			if err != nil {
				return err
			}

			if i == 0 {
				_, err = f.WriteString(template)
			}
			if cerr := f.Close(); err == nil {
				err = cerr
			}
			if err != nil {
				return err
			}
		}

	}
	return nil
}
