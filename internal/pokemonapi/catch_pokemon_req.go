package pokemonapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Pokemon(name string) (Pokemon, error) {
	if len(name) == 0 {
		return Pokemon{}, errors.New("need a pokemon name value")
	}

	enpoint := "/pokemon/"
	fullURL := baseURL + enpoint + name

	cacheHit, ok := c.cache.Get(fullURL)

	if ok {
		fmt.Println("Cache Hit")
		data := Pokemon{}

		if err := json.Unmarshal(cacheHit, &data); err != nil {
			return Pokemon{}, fmt.Errorf("failed to unmarshal json data: %w", err)
		}
		return data, nil
	}

	fmt.Println("Cache Miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, body)

	data := Pokemon{}

	if err := json.Unmarshal(body, &data); err != nil {
		return Pokemon{}, err
	}

	return data, nil
}
