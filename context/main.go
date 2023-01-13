package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type jokeResponse struct {
	Categories []string `json:"categories"`
	CreatedAt  string   `json:"created_at"`
	IconURL    string   `json:"icon_url"`
	ID         string   `json:"id"`
	UpdatedAt  string   `json:"updated_at"`
	URL        string   `json:"url"`
	Value      string   `json:"value"`
}

func fetchJoke(url string, c chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var joke jokeResponse

	err = json.Unmarshal(body, &joke)
	if err != nil {
		panic(err)
	}

	c <- joke.Value
	close(c)
}

func main() {
	now := time.Now()
	c := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	go fetchJoke("https://api.chucknorris.io/jokes/random", c)

	select {
	case <-ctx.Done():
		fmt.Println("Timed out")
		return
	case joke := <-c:
		fmt.Println(joke)
	}

	fmt.Println("Time taken:", time.Since(now))
}
