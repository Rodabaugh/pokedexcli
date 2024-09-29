package main

import (
	"fmt"
	"log"

	"github.com/rodabaugh/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	pokeapiClinet := pokeapi.NewClient()

	res, err := pokeapiClinet.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	return nil
}
