package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChuckNorrisJoke struct {
	Value string `json:"value"`
	err   error
}

func GetChuckNorrisJoke(ch chan<- ChuckNorrisJoke) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		ch <- ChuckNorrisJoke{err: err}
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	joke := ChuckNorrisJoke{}
	err = json.Unmarshal(body, &joke)
	if err != nil {
		ch <- ChuckNorrisJoke{err: err}
		return
	}

	ch <- joke
}

type DadJoke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	err       error
}

func GetDadJoke(ch chan<- DadJoke) {
	resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		ch <- DadJoke{err: err}
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	joke := DadJoke{}
	err = json.Unmarshal(body, &joke)
	if err != nil {
		ch <- DadJoke{err: err}
		return
	}

	ch <- joke
}

func main() {
	chuckNorrisCh := make(chan ChuckNorrisJoke)
	dadCh := make(chan DadJoke)

	go GetChuckNorrisJoke(chuckNorrisCh)
	go GetDadJoke(dadCh)

	chuckNorrisJoke := <-chuckNorrisCh
	if chuckNorrisJoke.err != nil {
		fmt.Println(chuckNorrisJoke.err)
		return
	}

	dadJoke := <-dadCh
	if dadJoke.err != nil {
		fmt.Println(dadJoke.err)
		return
	}

	fmt.Println(chuckNorrisJoke.Value)
	fmt.Println(dadJoke.Setup, dadJoke.Punchline)
}
