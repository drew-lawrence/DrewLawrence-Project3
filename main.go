package main

import (
	"fmt"
)

func main() {
	// Fetch all countries in Europe from the live API
	countries, err := fetchCountries("Asia")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print each country's common name
	for _, country := range countries {
		fmt.Println(country.Names.Common)
	}
}
