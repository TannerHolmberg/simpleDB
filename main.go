package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Read input from stdin
	in := bufio.NewScanner(os.Stdin)

	//For large input, increase the buffer size
	in.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	//Temp in-memory store for key-value pairs for validation and retrieval
	mem := make(map[string]string)

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
			if val, exists := mem[key]; exists {
				fmt.Println(val)
			}
		case "SET":
			mem[key] = value
		case "Invalid":
			continue
		}

	}

}
