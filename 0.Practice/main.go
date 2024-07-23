package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.foxnxx.com/")
	fmt.Println("Data received")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	data := resp.Body

	file, err := os.Create("data.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	scan := bufio.NewScanner(data)
	var i int
	for scan.Scan() {
		i++
		fmt.Println("Writing...", i)
		data := scan.Bytes()
		file.Write(data)
	}

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./data.html")
		return
	})

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
