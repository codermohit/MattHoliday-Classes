package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

/*
  This program uses concurrency to get urls , on running this program checkout the log timings and the time in which the request was fulfilled
  It uses context.WithTimeout() to cancel the requests which doesn't respond within the time limit
  use 'time go run .' : to get the time for which the program ran
*/

type result struct {
	url     string
	err     error
	latency time.Duration
}

func getUrl(ctx context.Context, url string, ch chan<- result) {
	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if resp, err := http.DefaultClient.Do(req); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}

}

func main() {
	results := make(chan result)
	list := []string{"https://google.com", "https://facebook.com", "https://msn.org", "https://bing.com", "https://wsj.com", "https://nytimes.com", "https://gobyexample.com"}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for _, url := range list {
		go getUrl(ctx, url, results)
	}

	for range list {
		r := <-results

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}
