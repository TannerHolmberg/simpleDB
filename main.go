package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	index := NewIndex()

	//Read input from stdin
	in := bufio.NewScanner(os.Stdin)

	//For large input, increase the buffer size
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
			indexValue, found := index.Get(key)
			if found {
				fmt.Println(indexValue)
			}
		case "SET":
			index.Set(key, value)
		case "Invalid":
			continue
		}

	}

}
