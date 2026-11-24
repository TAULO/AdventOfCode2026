package main

func main() {
	// TODO: read from os.Args and add .gitignore
	var err = createDay()

	if err != nil {
		panic(err)
	}
}

// to run any day, cd <day-to-run> && go run .
