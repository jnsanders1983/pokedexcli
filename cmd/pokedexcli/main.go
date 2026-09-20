package main

import (
	"time"

	"github.com/jnsanders1983/pokedexcli/internal/commands"
	"github.com/jnsanders1983/pokedexcli/internal/pokecache"
	"github.com/jnsanders1983/pokedexcli/repl"
)

func main() {
	cfg := commands.NewConfig(pokecache.NewCache(5 * time.Second))
	repl.Start(cfg)
}
