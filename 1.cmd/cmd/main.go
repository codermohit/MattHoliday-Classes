package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {
	fmt.Println(hello.SayHello(os.Args[1:]))
}
