package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// create the Fyne application
	myApp := app.New()

	// create a window with a title
	w := myApp.NewWindow("Country Explorer")

	var countries []Country
	// labels that show the selected country's details.
	// wrapping lets long values (like official names) use multiple lines.
	officialLabel := widget.NewLabel("")
	commonLabel := widget.NewLabel("")
	languagesLabel := widget.NewLabel("")
	landlockedLabel := widget.NewLabel("")
	capitalLabel := widget.NewLabel("")
	populationLabel := widget.NewLabel("")

	// flag image: starts empty, filled in when a country is clicked
	flagImage := canvas.NewImageFromResource(nil)
	flagImage.FillMode = canvas.ImageFillContain // keep the flag's proportions
	flagImage.SetMinSize(fyne.NewSize(240, 160)) // without this it can render at 0x0
	for _, l := range []*widget.Label{officialLabel, commonLabel, languagesLabel, capitalLabel} {
		l.Wrapping = fyne.TextWrapWord
	}

	// a form gives each value a clear label on the left
	details := widget.NewForm(
		widget.NewFormItem("Flag", flagImage),
		widget.NewFormItem("Official name", officialLabel),
		widget.NewFormItem("Common name(s)", commonLabel),
		widget.NewFormItem("Language(s)", languagesLabel),
		widget.NewFormItem("Landlocked", landlockedLabel),
		widget.NewFormItem("Capital(s)", capitalLabel),
		widget.NewFormItem("Population", populationLabel),
	)

	// card around the details: big title + a hint until a country is picked
	card := widget.NewCard("No country selected", "Choose a country from the list to see its details", details)
	details.Hide() // hide the empty form until something is clicked

	// creates list of country names
	list := widget.NewList(
		// how many rows in the list
		func() int {
			return len(countries)
		},
		// single row
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		// fill row i with country's common name
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(countries[id].Names.Common)
		},
	)

	// when a country in the list is clicked, show its details
	list.OnSelected = func(i widget.ListItemID) {
		c := countries[i]

		// show the country's name as the card title and reveal the details
		card.SetTitle(c.Names.Common)
		card.SetSubTitle("")
		details.Show()

		officialLabel.SetText(c.Names.Official)

		// common name plus any alternate names
		names := append([]string{c.Names.Common}, c.Names.Alternates...)
		commonLabel.SetText(strings.Join(names, ", "))

		// collect each language's English name
		var langs []string
		for _, lang := range c.Languages {
			langs = append(langs, lang.Name)
		}
		languagesLabel.SetText(strings.Join(langs, ", "))

		// turn true/false into Yes/No for readability
		if c.Landlocked {
			landlockedLabel.SetText("Yes")
		} else {
			landlockedLabel.SetText("No")
		}

		// collect capital city names
		var caps []string
		for _, cap := range c.Capitals {
			caps = append(caps, cap.Name)
		}
		capitalLabel.SetText(strings.Join(caps, ", "))

		populationLabel.SetText(formatWithCommas(c.Population))
		// download and show the flag
		flagBytes, err := fetchFlag(c.Flag.URLPng)
		if err != nil {
			fmt.Println(err)
			flagImage.Resource = nil
		} else {
			flagImage.Resource = fyne.NewStaticResource(c.Names.Common+".png", flagBytes)
		}
		flagImage.Refresh()
	}

	// Claude helped with this regionSelect widget
	regionSelect := widget.NewSelect(
		[]string{"Africa", "Americas", "Asia", "Europe", "Oceania"},
		func(region string) {
			fetched, err := fetchCountries(region)
			if err != nil {
				fmt.Println(err)
				return
			}

			countries = fetched
			list.UnselectAll()

			// reset the details side when switching regions
			card.SetTitle("No country selected")
			card.SetSubTitle("Choose a country from the list to see its details")
			details.Hide()
			list.Refresh()
		},
	)
	regionSelect.PlaceHolder = "Choose a region"

	// list on the left (30%), details on the right
	// Claude suggestion
	split := container.NewHSplit(list, container.NewVScroll(card))
	split.Offset = 0.3

	// big title and instructions above the dropdown
	title := widget.NewRichTextFromMarkdown("# Country Explorer")
	instructions := widget.NewLabel("Choose a region below, then click a country to see its details.")
	top := container.NewVBox(title, instructions, regionSelect)

	// dropdown on top, split view fills the rest
	w.SetContent(container.NewBorder(top, nil, nil, nil, split))
	w.Resize(fyne.NewSize(1200, 800))
	w.ShowAndRun()
}

// formatWithCommas turns 2335930 into "2,335,930"
// Claude suggestion
func formatWithCommas(n int64) string {
	s := strconv.FormatInt(n, 10)
	// Work backwards from the end, inserting a comma every 3 digits
	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	return s
}
