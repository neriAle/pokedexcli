package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		cleanText := cleanInput(reader.Text())
		if len(cleanText) < 1 {
			continue
		}

		cmd, ok := getCommands()[cleanText[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	return strings.Fields(lowercaseText)
}