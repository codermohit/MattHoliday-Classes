package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

const url = "http://jsonplaceholder.typicode.com/"

type todo struct {
	UserID    int    `json:"UserID"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>
`

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(url + r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return //must return
	}

	//if resp is not closed, then the websocket keeps running
	defer resp.Body.Close()

/*
	body, err := io.ReadAll(resp.Body)
	if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
	}
*/

	var item todo

	err = json.NewDecoder(resp.Body).Decode(&item)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  tmpl := template.New("mine")

  tmpl.Parse(form)
  tmpl.Execute(w, item)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
