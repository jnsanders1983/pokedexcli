package main

func main() {
	cfg := &config{
		commands: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Display help information",
				callback:    commandHelp,
			},
			"map": {
				name: "map",
				description: `Displays the names of 20 location areas in the Pokemon world.
				 Each subsequent call to the map command will display the 
				 next 20 location areas`,
				callback: commandMap,
			},
			"mapb": {
				name: "mapb",
				description: `Displays the names of the previous 20 location areas in the 
				 Pokemon world. Each subsequent call to the mapb command will display the 
				 previous 20 location areas`,
				callback: commandMapBack,
			},
		},
	}
	startRepl(cfg)
}
