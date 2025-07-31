package repl

import(
	"bufio"
	"fmt"
	"github.com/neriAle/pokedexcli/internal/pokecache"
	"os"
	"strings"
	"time"
)

type Config struct {
	Next_area 		string
	Previous_area 	string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, *pokecache.Cache, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func StartREPL() {
	reader := bufio.NewScanner(os.Stdin)
	conf := Config{}
	cache := pokecache.NewCache(15 * time.Second)
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

		err := cmd.callback(&conf, cache, cleanText[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	return strings.Fields(lowercaseText)
}