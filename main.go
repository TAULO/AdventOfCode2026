package main

func main() {
	// TODO: read from os.Args and add .gitignore
	var err = createDay(1, "Foobar")

	if err != nil {
		panic(err)
	}
}

// to run any day, cd <day-to-run> && go run .
