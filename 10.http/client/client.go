package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "http://jsonplaceholder.typicode.com"

type todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	resp, err := http.Get(url + "/todos/1")

	if err != nil {
		log.Fatal(err)
	}

	//if resp is not closed, then the websocket keeps running
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		var item todo

		if err := json.Unmarshal(body, &item); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v\n", item)
	}
}
