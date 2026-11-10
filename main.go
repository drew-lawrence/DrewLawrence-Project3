package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	os.ReadFile("europe.json")

	// read entire JSON file as raw bytes
	bytes, err := os.ReadFile("europe.json")
	if err != nil {
		// error handling
		fmt.Println(err)
	}
	// create am empty APIResponse for the JSON to be loaded into
	var response APIResponse

	// convert raw JSON bytes into structs
	err = json.Unmarshal(bytes, &response)

	// loop through every country in data.objects and print its common name
	for _, country := range response.Data.Objects {
		fmt.Println(country.Names.Common)
	}

}
