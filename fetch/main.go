package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var joke jokeResponse

	err = json.Unmarshal(body, &joke)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(joke.Value)
}
