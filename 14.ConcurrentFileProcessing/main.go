package main

import (
	"fmt"
	"runtime"
)

func main() {
	workers := runtime.GOMAXPROCS(0)
	fmt.Println(workers)
}
