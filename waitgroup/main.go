package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = map[string]string{
	"google":   "https://google.com",
	"yahoo":    "https://yahoo.com",
	"bing":     "https://bing.com",
	"notasite": "https://notasite.com",
}

func checkSiteStatus(wg *sync.WaitGroup, site string, url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s: DOWN\n", site)
		return
	}
	fmt.Printf("%s: %s\n", site, res.Status)
}

func main() {
	var wg sync.WaitGroup
	for site, url := range urls {
		wg.Add(1)
		go checkSiteStatus(&wg, site, url)
	}
	wg.Wait()
}
