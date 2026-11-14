# Country Explorer - Project 3

By: Drew Lawrence

Project Description: A Fyne desktop app that uses the REST Countries API to show information about countries in a chosen region.

## Setup

1. Get a free API key from https://restcountries.com
2. In the project folder, create a file named `api_secrets.go` containing:
   
   go
   package main

   var api_key = "your-key-here"

   This file is listed in `.gitignore`, so the key is never uploaded to GitHub.
3. Fyne requires a 64-bit C compiler (gcc). On Windows I used MSYS2's UCRT64 gcc.
4. Run `go mod tidy`, then `go run .` The first build takes a few minutes.

## How to use

1. Choose a region from the dropdown at the top (Africa, Americas, Asia, Europe, or Oceania).
2. The list on the left fills with every country in that region.
3. Click a country to see its details on the right: flag, official name, common name(s), languages, whether it is landlocked, capital(s), and population.
4. Choosing a different region resets the details panel.

## Decisions

- Region vs. membership: I chose region, because every country belongs to just one region. A country could either have multiple or no membership at all. 
- Name shown in the list: the common name (e.g. "Albania" instead of "Republic of Albania"), because it's shorter and what most people recognize. The official name is shown in the details.
- Widget for the countries: a Fyne List, because I only needed to display country names (a single column), and the Fyne list has built-in click selection for showing details.
- Common name(s): the details show the common name plus any alternate names from the API (e.g. Albania, Shqipëri, Shqipnia).
- Languages: the API lists a country's languages but doesn't mark which are official, so all listed languages are shown.
- Extra pieces -> Capital(s) and Population: Capital because it's a popular thing to know about a country, and the API handles multiple capitals. Population because it shows size at a glance, and it is current data as the API updates it every few hours.
- Getting every country: the API returns only 25 results by default, so the request uses `limit=100` so larger regions like Africa and Europe aren't cut off.
- Window size: 1200x800, which fits within a 1900x1080 display.
- Sort order: countries come back sorted alphabetically from the API. Names starting with accented letters (like Åland Islands) appear at the end.

## Explaining it to someone outside the class

I showed this to my roommate Dane who is not in the class. They picked a region and clicked through several countries on their own.
I did not have to explain any of the features, due to the instructions and labels that are displayed on the GUI. 

