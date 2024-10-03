package pokemonapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	cacheHit, ok := c.cache.Get(fullURL)

	if ok {
		fmt.Println("Cache hit")
		var data LocationAreasResponse

		if err := json.Unmarshal(cacheHit, &data); err != nil {
			return LocationAreasResponse{}, fmt.Errorf("failed to unmarshal json data: %w", err)
		}
		return data, nil
	}

	fmt.Println("Cache Miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, body)

	var data LocationAreasResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return LocationAreasResponse{}, fmt.Errorf("failed to unmarshal json data: %w", err)
	}

	return data, nil
}
