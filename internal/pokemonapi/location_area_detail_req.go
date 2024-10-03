package pokemonapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationAreaDetail(area_name string) (LocationAreaDetail, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint + "/" + area_name

	cacheHit, ok := c.cache.Get(fullURL)

	if ok {
		fmt.Println("Cache Hit")
		var data LocationAreaDetail

		if err := json.Unmarshal(cacheHit, &data); err != nil {
			return LocationAreaDetail{}, fmt.Errorf("failed to unmarshal json data: %w", err)
		}

		return data, nil
	}

	fmt.Println("Cache Miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaDetail{}, fmt.Errorf("bad status code: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	c.cache.Add(fullURL, body)

	var data LocationAreaDetail

	if err := json.Unmarshal(body, &data); err != nil {
		return LocationAreaDetail{}, fmt.Errorf("failed to unmarshal json data: %w", err)
	}

	return data, nil
}
