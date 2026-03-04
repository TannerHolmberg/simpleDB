package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Read input from stdin
	in := bufio.NewScanner(os.Stdin)

	//For large input, increase the buffer size
	in.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	for in.Scan() {
		//Trim leading and trailing whitespace
		line := strings.TrimSpace(in.Text())

		//Skip empty lines
		if line == "" {
			continue
		}

		//Check for exit command (case-insensitive)
		if strings.EqualFold(line, "EXIT") {
			return
		}

		//Error handling
		if err := in.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
	}

}
