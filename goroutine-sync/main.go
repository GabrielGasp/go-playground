package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type ChuckNorrisJoke struct {
	Value string `json:"value"`
}

func GetChuckNorrisJoke(wg *sync.WaitGroup, ch chan<- string, errCh chan<- error) {
	defer wg.Done()

	// errCh <- fmt.Errorf("this is an error")

	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		errCh <- err
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	joke := ChuckNorrisJoke{}
	err = json.Unmarshal(body, &joke)
	if err != nil {
		errCh <- err
		return
	}

	ch <- joke.Value
}

type DadJoke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func GetDadJoke(wg *sync.WaitGroup, ch chan<- string, errCh chan<- error) {
	defer wg.Done()

	// errCh <- fmt.Errorf("this is an error")

	resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		errCh <- err
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	joke := DadJoke{}
	err = json.Unmarshal(body, &joke)
	if err != nil {
		errCh <- err
		return
	}

	ch <- joke.Setup + " " + joke.Punchline
}

func main() {
	wg := sync.WaitGroup{}
	chuckNorrisCh := make(chan string, 1)
	dadCh := make(chan string, 1)
	errCh := make(chan error, 1)

	wg.Add(2)

	go GetChuckNorrisJoke(&wg, chuckNorrisCh, errCh)
	go GetDadJoke(&wg, dadCh, errCh)

	go func() {
		wg.Wait()
		close(chuckNorrisCh)
		close(dadCh)
		close(errCh)
	}()

	for err := range errCh {
		fmt.Println(err)
		return
	}

	for joke := range chuckNorrisCh {
		fmt.Println(joke)
	}

	for joke := range dadCh {
		fmt.Println(joke)
	}

	fmt.Println("Done!")
}
