package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// Create the Fyne application
	myApp := app.New()

	// Create a window with a title
	w := myApp.NewWindow("Country Explorer")

	// Put a simple label in the window as a test
	w.SetContent(widget.NewLabel("Hello"))

	// Set the starting size (stays under the 1900x1080 limit)
	w.Resize(fyne.NewSize(1200, 800))

	// Show the window and keep the program running until it's closed
	w.ShowAndRun()

	/*
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
	*/
}
