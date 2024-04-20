package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type ChuckNorrisJoke struct {
	Value string `json:"value"`
	err   error
}

func GetChuckNorrisJoke() (string, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	joke := ChuckNorrisJoke{}
	err = json.Unmarshal(body, &joke)
	if err != nil {
		return "", err
	}

	return joke.Value, nil
}

type DadJoke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	err       error
}

func GetDadJoke() (string, error) {
	resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	joke := DadJoke{}
	err = json.Unmarshal(body, &joke)
	if err != nil {
		return "", err
	}

	return joke.Setup + " " + joke.Punchline, nil
}

func main() {
	var chuckNorrisJoke string
	var dadJoke string

	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		var err error
		chuckNorrisJoke, err = GetChuckNorrisJoke()
		return err
	})

	g.Go(func() error {
		var err error
		dadJoke, err = GetDadJoke()
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(chuckNorrisJoke)
	fmt.Println(dadJoke)
}
