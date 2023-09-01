package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	t := time.Now()

	f, err := os.Open("icon_urls.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	g, _ := errgroup.WithContext(context.Background())
	g.SetLimit(100)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		icon := scanner.Text()
		g.Go(func() error {
			return downloadIcon(icon)
		})
	}

	if err := g.Wait(); err != nil {
		panic(err)
	}

	fmt.Printf("Done in %v\n", time.Since(t))
}

func downloadIcon(icon string) error {
	f, err := os.Create("icons/" + icon)
	if err != nil {
		return err
	}

	defer f.Close()

	url := "https://d1fojj4wte942r.cloudfront.net/e-games/" + icon
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error downloading %s\n", icon)
		return err
	}

	defer res.Body.Close()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		return err
	}

	return nil
}
