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

	// errCh <- fmt.Errorf("chuck norris error")

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

	// errCh <- fmt.Errorf("dad error")

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
	errCh := make(chan error, 2)

	wg.Add(2)

	go GetChuckNorrisJoke(&wg, chuckNorrisCh, errCh)
	go GetDadJoke(&wg, dadCh, errCh)

	go func() {
		wg.Wait()
		close(chuckNorrisCh)
		close(dadCh)
		close(errCh)
	}()

	errs := make([]error, 0, 2)
	for err := range errCh {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		fmt.Println("Done with error!")
		return
	}

	chuckNorrisJoke := <-chuckNorrisCh
	fmt.Println(chuckNorrisJoke)

	dadJoke := <-dadCh
	fmt.Println(dadJoke)

	fmt.Println("Done with jokes!")
}
