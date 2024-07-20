package main

import (
	"fmt"
	"time"
)

func main() {
	const tickRate = 2 * time.Second
	ticker := time.NewTicker(tickRate).C
	stopper := time.After(5 * tickRate)
loop:
	for {
		select {
		case <-ticker:
			fmt.Println("tick")

		case <-stopper:
			break loop

		}
	}
}
