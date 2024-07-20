package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	stopper := time.After(12 * time.Second)
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		fmt.Println(i)
		go func(i int, ch chan int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}

	for {
		select {
		case m0 := <-chans[0]:
			fmt.Println("received", m0)

		case m1 := <-chans[1]:
			fmt.Println("received", m1)

		case <-stopper:
			fmt.Println("timeout , stopping after 12 secs")
			os.Exit(0)
		}
	}
}

