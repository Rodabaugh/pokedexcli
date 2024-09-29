package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationURL = res.Next
	cfg.previousLocationURL = res.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocationURL == nil {
		return errors.New("already on the first page")
	}

	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationURL)
	if err != nil {
		return err
	}

	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationURL = res.Next
	cfg.previousLocationURL = res.Previous

	return nil
}
