package main

/*
These structs were created by Claude, but I understand how they are created and how they work.
*/
type APIResponse struct {
	Data struct {
		Objects []Country `json:"objects"`
	} `json:"data"`
}

type Country struct {
	Names struct {
		Common     string   `json:"common"`
		Official   string   `json:"official"`
		Alternates []string `json:"alternates"`
	} `json:"names"`
	Capitals []struct {
		Name string `json:"name"`
	} `json:"capitals"`
	Flag struct {
		URLPng string `json:"url_png"`
	} `json:"flag"`
	Landlocked bool `json:"landlocked"`
	Languages  []struct {
		Name       string `json:"name"`
		NativeName string `json:"native_name"`
	} `json:"languages"`
	Population int64 `json:"population"`
}
