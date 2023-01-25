package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fetchJoke("https://api.chucknorris.io/jokes/random")
}

type jokeResponse struct {
	Categories []string `json:"categories"`
	CreatedAt  string   `json:"created_at"`
	IconURL    string   `json:"icon_url"`
	ID         string   `json:"id"`
	UpdatedAt  string   `json:"updated_at"`
	URL        string   `json:"url"`
	Value      string   `json:"value"`
}

func fetchJoke(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: ", resp.StatusCode)
		return
	}

	var joke jokeResponse

	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(joke.Value)
}
