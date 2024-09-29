package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("one pokemon must be provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s", pokemonName)
	time.Sleep(time.Second / 4)
	fmt.Print(".")
	time.Sleep(time.Second / 4)
	fmt.Print(".")
	time.Sleep(time.Second / 4)
	fmt.Print(".")
	time.Sleep(time.Second / 4)
	fmt.Print("\n")

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
