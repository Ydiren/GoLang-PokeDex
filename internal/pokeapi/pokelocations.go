package pokeapi

import "encoding/json"

type PokeLocations struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (locations *PokeLocations) parseData(data []byte) error {
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return err
	}

	return nil
}
