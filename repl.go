package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		cleanText := cleanInput(reader.Text())
		if len(cleanText) > 0 {
			fmt.Printf("Your command was: %s\n", cleanText[0])
		}
	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	return strings.Fields(lowercaseText)
}