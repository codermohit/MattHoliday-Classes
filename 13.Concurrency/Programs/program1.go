package main

import (
	"log"
	"net/http"
	"time"
)

/*
  This program uses concurrency to get urls , on running this program checkout the log timings and the time in which the request was fulfilled
  use 'time go run .' : to get the time for which the program ran
*/

type result struct {
	url     string
	err     error
	latency time.Duration
}

func getUrl(url string, ch chan<- result) {
	start := time.Now()
	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}

}

func main() {
	results := make(chan result)
	list := []string{"https://google.com", "https://facebook.com", "https://msn.org", "https://bing.com", "https://wsj.com", "https://nytimes.com"}

	for _, url := range list {
		go getUrl(url, results)
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

