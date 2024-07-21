package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex> ")
		scanner.Scan()
		cmd := scanner.Text()
		cleaned := cleanInput(cmd)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]

		switch command {
		case "exit":
			os.Exit(1)
		}

		fmt.Println(cleaned)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
