package main

import (
	"strings"
)

// parseInput parses a single line of user input into a command, key, and optional value.
// It supports the commands:
//   - EXIT
//   - GET <key>
//   - SET <key> <value...>
//
// It returns ok=false for empty or invalid input; otherwise ok=true with the normalized
// command string (e.g., "GET", "SET", "EXIT"). For SET, the value may contain spaces and
// is reconstructed from the remaining fields.
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
		return "Invalid", "", "", false
	}

	command := strings.ToUpper(parts[0])
	switch command {
	case "GET":
		if len(parts) < 2 {
			return "Invalid", "", "", false
		}
		return "GET", parts[1], "", true
	case "SET":
		if len(parts) < 3 {
			return "Invalid", "", "", false
		}
		return "SET", parts[1], strings.Join(parts[2:], " "), true
	default:
		return "Invalid", "", "", false
	}
}
