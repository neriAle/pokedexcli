package main

import(
	"fmt"
	"os"
	"github.com/neriAle/pokedexcli/pokeapi"
)

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("usage:\n")

	for _, v := range getCommands() {
		fmt.Printf("%s:\t%s\n", v.name, v.description)
	}
	return nil
}

func commandMap(c *Config) error {
	areas, prev, next, err := pokeapi.Get_location_areas(c.Previous_area, c.Next_area)
	if err != nil {
		return err
	}

	c.Previous_area = prev
	c.Next_area = next

	for _, a := range areas {
		fmt.Println(a)
	}
	return nil
}

func commandMapb(c *Config) error {
	// If we are on the first page, just print it and exit
	if c.Previous_area == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	areas, prev, next, err := pokeapi.Get_location_areas(c.Next_area, c.Previous_area)
	if err != nil {
		return err
	}

	c.Previous_area = prev
	c.Next_area = next

	for _, a := range areas {
		fmt.Println(a)
	}
	return nil
}