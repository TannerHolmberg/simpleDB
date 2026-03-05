package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s, err := openStore("data.db")
	if err != nil {
		return
	}
	defer s.file.Close()

	//Read input from stdin
	in := bufio.NewScanner(os.Stdin)

	// For large input, increase the buffer size
	in.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	for in.Scan() {
		//Trim leading and trailing whitespace
		line := in.Text()

		//Parse the input line into command, key, value, and validity
		cmd, key, value, ok := parseInput(line)

		if !ok {
			continue
		}

		if cmd == "EXIT" {
			break
		}

		switch cmd {
		case "GET":
			Value, found := s.Get(key)
			if found {
				fmt.Println(Value)
			} else {
				fmt.Println()
			}
		case "SET":
			_ = s.Set(key, value)
		case "Invalid":
			continue
		}

	}

	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
