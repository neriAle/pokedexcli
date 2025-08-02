package repl

import(
	"encoding/json"
	"fmt"
	"github.com/neriAle/pokedexcli/internal/pokeapi"
	"github.com/neriAle/pokedexcli/internal/pokecache"
	"math/rand"
	"os"
)

func commandExit(c *Config, cache *pokecache.Cache, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config, cache *pokecache.Cache, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("usage:\n\n")

	for _, v := range getCommands() {
		fmt.Printf("%-20s: %s\n", v.name, v.description)
	}
	return nil
}

func commandMap(c *Config, cache *pokecache.Cache, args ...string) error {
	// areas, prev, next, err := pokeapi.Get_location_areas(c.Previous_area, c.Next_area)
	var url string
	if c.Next_area == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = c.Next_area
	}

	var data []byte
	// Check if the value at this url is already in the cache
	value, ok := cache.Get(url)
	if ok {
		data = value
	} else {
		val, err := pokeapi.Get_api_data(url)
		if err != nil {
			return err
		}
		data = val
		cache.Add(url, data)
	}

	// Unmarshal the data into a Location_Areas struct
	var locations = pokeapi.Location_Areas{}
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return fmt.Errorf("error unmarshalling the location areas: %w", err)
	}

	// Update the next and previous location of the config struct
	if locations.Next != nil {
		c.Next_area = *locations.Next
	}
	if locations.Previous != nil {
		c.Previous_area = *locations.Previous
	} else {
		c.Previous_area = ""
	}

	// Print the results
	for _, a := range locations.Results {
		fmt.Println(a.Name)
	}
	return nil
}

func commandMapb(c *Config, cache *pokecache.Cache, args ...string) error {
	// If we are on the first page, just print it and exit
	if c.Previous_area == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := c.Previous_area

	var data []byte
	// Check if the value at this url is already in the cache
	value, ok := cache.Get(url)
	if ok {
		data = value
	} else {
		val, err := pokeapi.Get_api_data(url)
		if err != nil {
			return err
		}
		data = val
		cache.Add(url, data)
	}

	// Unmarshal the data into a Location_Areas struct
	var locations = pokeapi.Location_Areas{}
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return fmt.Errorf("error unmarshalling the location areas: %w", err)
	}

	// Update the next and previous location of the config struct
	if locations.Next != nil {
		c.Next_area = *locations.Next
	}
	if locations.Previous != nil {
		c.Previous_area = *locations.Previous
	} else {
		c.Previous_area = ""
	}

	// Print the results
	for _, a := range locations.Results {
		fmt.Println(a.Name)
	}
	return nil
}

func commandExplore(c *Config, cache *pokecache.Cache, args ...string) error {
	// If no argument has been passed, ask the user to input one
	if len(args) < 1 {
		fmt.Println("Please insert an area to explore")
		return nil
	}
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]

	fmt.Printf("Exploring %s...\n", args[0])

	var data []byte
	// Check if the value at this url is already in the cache
	value, ok := cache.Get(url)
	if ok {
		data = value
	} else {
		val, err := pokeapi.Get_api_data(url)
		if err != nil {
			return fmt.Errorf("Not a valid location area, %w", err)
		}
		data = val
		cache.Add(url, data)
	}

	// Unmarshal the data into a Location_Areas struct
	var location = pokeapi.Location_Details{}
	err := json.Unmarshal(data, &location)
	if err != nil {
		return fmt.Errorf("error unmarshalling the location areas: %w", err)
	}

	// If no pokemon can be encountered in this area, print a message
	if len(location.PokemonEncounters) < 1 {
		fmt.Println("No pokemon can be encountered in this area")
		return nil
	}

	fmt.Println("Found Pokemon:")

	// Print the results
	for _, pe := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pe.Pokemon.Name)
	}
	return nil
}

func commandCatch(c *Config, cache *pokecache.Cache, args ...string) error {
	// If no argument has been passed, ask the user to input one
	if len(args) < 1 {
		fmt.Println("Please insert the name of a pokemon to catch")
		return nil
	}
	pokemon_name := args[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon_name

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon_name)

	var data []byte
	// Check if the value at this url is already in the cache
	value, ok := cache.Get(url)
	if ok {
		data = value
	} else {
		val, err := pokeapi.Get_api_data(url)
		if err != nil {
			return fmt.Errorf("Not a valid pokemon name, %w", err)
		}
		data = val
		cache.Add(url, data)
	}

	// Unmarshal data into a Pokemon struct
	var pkmn = pokeapi.Pokemon{}
	err := json.Unmarshal(data, &pkmn)
	if err != nil {
		return fmt.Errorf("error unmarshalling the location areas: %w", err)
	}

	baseExp := pkmn.BaseExperience
	if baseExp > 350 {
		baseExp = 350
	}
	chance := rand.Intn(400)

	// If the random chance is higher than the base experience of the pokemon, catch it and add it to the pokedex
	if chance > baseExp {
		fmt.Printf("%s was caught!\n", pokemon_name)
		fmt.Printf("You may now inspect it with the inspect command.\n")
		c.Pokedex[pokemon_name] = pkmn
	} else {
		fmt.Printf("%s escaped!\n", pokemon_name)
	}

	return nil
}

func commandInspect(c *Config, cache *pokecache.Cache, args ...string) error {
	// If no argument has been passed, ask the user to input one
	if len(args) < 1 {
		fmt.Println("Please insert the name of a pokemon to inspect")
		return nil
	}
	pokemon_name := args[0]

	pkmn, caught := c.Pokedex[pokemon_name]
	if !caught {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	details := fmt.Sprintf("\nName: %s\nHeight: %d\nWeight: %d\n", pkmn.Name, pkmn.Height, pkmn.Weight)
	stats := fmt.Sprintf("Stats:\n\t- hp: %d\n\t- attack: %d\n\t- defense: %d\n\t- special-attack: %d\n\t- special-defense: %d\n\t- speed: %d\n", 
		pkmn.Stats[0].BaseStat, pkmn.Stats[1].BaseStat, pkmn.Stats[2].BaseStat, pkmn.Stats[3].BaseStat, pkmn.Stats[4].BaseStat, pkmn.Stats[5].BaseStat)
	types := "Types:\n"
	for i := range pkmn.Types {
		types = types + fmt.Sprintf("\t- %s\n", pkmn.Types[i].Type.Name)
	}

	final_string := details + stats + types
	fmt.Println(final_string)
	return nil
}

func commandPokedex(c *Config, cache *pokecache.Cache, args ...string) error {
	if len(c.Pokedex) < 1 {
		fmt.Println("You have not caught any pokemon yet!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for k, _ := range c.Pokedex {
		fmt.Printf("\t- %s\n", k)
	}

	return nil
}