package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// fetchCountries asks the REST Countries API for every country in a region and returns them as a slice of Country structs.
func fetchCountries(region string) ([]Country, error) {
	// Build the request URL with the chosen region.
	// limit=100 so we get every country, not just the first 25.
	// response_fields keeps the response small by only asking for what we display.
	url := fmt.Sprintf(
		"https://api.restcountries.com/countries/v5/region/%s?limit=100&response_fields=names.common,names.official,names.alternates,languages,flag.url_png,landlocked,capitals,population",
		region,
	)

	// Create the GET request (nil = no request body)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Attach the API key from api_secrets.go as a bearer token
	req.Header.Set("Authorization", "Bearer "+api_key)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// Close the connection when this function finishes
	defer resp.Body.Close()

	// Read the whole response body into bytes
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// If the API returned an error (bad key, bad region, etc.), report it
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	// Convert the JSON into our structs
	var response APIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Hand back just the list of countries
	return response.Data.Objects, nil
}
