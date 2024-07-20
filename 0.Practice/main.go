package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Result struct {
	url     string
	err     error
	latency time.Duration
}

func getUrl(ctx context.Context, url string, ch chan<- Result) {
	start := time.Now()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		ch <- Result{url, err, 0}
	}

	if resp, err := http.DefaultClient.Do(req); err != nil {
		ch <- Result{url, err, 0}
	} else {
		ch <- Result{url, nil, time.Since(start).Round(time.Millisecond)}
		resp.Body.Close()
	}
}

func main() {

	list := []string{"https://google.com", "https://facebook.com", "https://msn.org", "https://bing.com", "https://wsj.com", "https://nytimes.com", "https://gobyexample.com"}

	resultChan := make(chan Result)

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	for _, url := range list {
		go getUrl(ctx, url, resultChan)
	}

	for range len(list) {
		result := <-resultChan

		if result.err != nil {
			log.Printf("%-20s %s\n", result.url, result.err.Error())
		} else {
			log.Printf("%-20s , %s", result.url, result.latency)
		}
	}
}
