package main

import (
	"encoding/json"
	"fmt"
)

/* Program to show the use of struct tags */

type Response struct {
	Name  int      `json:"page"`
	Words []string `json:"words,omitempty"`
}

func main() {
	r1 := Response{1, []string{"up", "down", "right"}}
	r2 := Response{Name: 2}

	j1, _ := json.Marshal(r1)
	j2, _ := json.Marshal(r2)

	fmt.Println(string(j1), string(j2))

	var jr1, jr2 Response

	_ = json.Unmarshal(j1, &jr1)
	_ = json.Unmarshal(j2, &jr2)

	fmt.Printf("%#v %#v\n", jr1, jr2)
}
