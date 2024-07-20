//generator - filter program to filter prime numbers upto a limit

package main

import "fmt"

func generator(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		ch <- i
	}

	close(ch)
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		fmt.Println(i)
		if i%prime != 0 {
			dst <- i
		}
	}

	close(dst)
}

func sieve(limit int) {
	ch := make(chan int)
	go generator(limit, ch)

	for {
		prime, ok := <-ch

		if !ok {
			break
		}

		ch1 := make(chan int)

		go filter(ch, ch1, prime)

		ch = ch1

		fmt.Print(prime, " ")
	}
	fmt.Println()
}

func main() {
	sieve(4)
}
