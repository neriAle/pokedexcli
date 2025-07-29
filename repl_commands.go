package main

import(
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("usage:\n")

	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}