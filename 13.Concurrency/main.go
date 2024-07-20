package main

import (
	"context"
	"log"
	"net/http"
	"runtime"
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

func getUrl(ctx context.Context, url string, ch chan<- result) {
	var r result

	start := time.Now()
	ticker := time.NewTicker(1 * time.Second).C
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req); err != nil {
		r = result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		r = result{url, nil, t}
		resp.Body.Close()
	}

	for {
		select {
		case ch <- r:
			return
		case <-ticker:
			log.Println("tick ", r)
		}
	}

}

func first(ctx context.Context, urls []string) (*result, error) {
	/*
	  In first line make(chan result) , we have no buffer to contain the result from the getUrl(), hence it would block the case where it has
	  to write on the results channel. This leads to a memory leak.
	  For make(chan result, len(urls)), we have a buffer to contain the result from the getUrl(), so as soon as the result is written on the
	  results channel , the Goroutine will terminate.
	*/
	//results := make(chan result)
	results := make(chan result, len(urls))
	ctx, cancel := context.WithCancel(ctx)

	//we have deferred the cancel here, as soon as we return from the function it will be called
	//so as soon as the first result is sent on the results channel by getUrl() and received in the first()
	//and return is executed , then the remaining requests will be cancelled
	defer cancel()

	for _, url := range urls {
		go getUrl(ctx, url, results)
	}

	select {
	case r := <-results:
		return &r, nil

		//it is important to handle the case when the context may get done, in this case before getUrl() receives any response
		//this may happen when the parent context may have a timeout on it.
	case <-ctx.Done():
		return nil, ctx.Err()

	}

}

func main() {
	list := []string{"https://google.com", "https://facebook.com", "https://msn.org", "https://bing.com", "https://wsj.com", "https://nytimes.com", "https://gobyexample.com"}

	r, _ := first(context.Background(), list)

	if r.err != nil {
		log.Printf("%-20s %s\n", r.url, r.err)
	} else {
		log.Printf("%-20s %s\n", r.url, r.latency)
	}

	time.Sleep(5 * time.Second)
	log.Println("quit anyway, ", runtime.NumGoroutine(), " still running")

}
