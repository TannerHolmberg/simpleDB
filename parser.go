package main

import (
	"strings"
)

func parseInput(input string) (Command string, key string, value string, ok bool) {
	//Trim leading and trailing whitespace
	line := strings.TrimSpace(input)
	//Check for exit command (case-insensitive)
	if strings.EqualFold(line, "EXIT") {
		return "EXIT", "", "", true
	}

	//Split the line into parts
	parts := strings.Fields(line)

	if len(parts) == 0 {
		return "Invalid", "", "", true
	}

	command := parts[0]
	switch command {
	case "GET":
		if len(parts) < 2 {
			return "Invalid", "", "", true
		}
		return "GET", parts[1], "", true
	case "SET":
		if len(parts) < 3 {
			return "Invalid", "", "", true
		}
		return "SET", parts[1], strings.Join(parts[2:], " "), true
	default:
		return "Invalid", "", "", true
	}
}
