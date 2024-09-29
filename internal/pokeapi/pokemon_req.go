package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	full_url := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(full_url)
	if ok {
		// the cache has hit
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", full_url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		if res.StatusCode == 404 {
			return Pokemon{}, fmt.Errorf("invalid pokemon: %s", pokemonName)
		}
		return Pokemon{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(full_url, data)

	return pokemon, nil
}
