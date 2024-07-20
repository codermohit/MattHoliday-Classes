package main

import (
	"fmt"
	"log"
	"net/http"
)

type nextCh chan int

func (ch nextCh) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h3>You id is : %d</h3>", <-ch)
}

func counter(ch chan<- int) {
	for i := 1; ; i++ {
		ch <- i
	}
}

func main() {
	var nextID nextCh = make(chan int)
	go counter(nextID)

	http.HandleFunc("/", nextID.handler)

	log.Fatal(http.ListenAndServe(":4000", nil))
}
