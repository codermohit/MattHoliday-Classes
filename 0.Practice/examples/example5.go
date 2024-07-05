package main

import (
	"fmt"
	"io"
	"os"
)

/*
  This Example shows the Method and Interface usage
  ByteCounter type implements the `Write(p []byte) (int, error)`, which implicitly implies that ByteCounter implements Writer Interface and will be accepted by methods that require a writer
*/

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*b += ByteCounter(l)

	f2, _ := os.Create("dfe.txt")
	s := fmt.Sprintf("%s", p)
	io.WriteString(f2, s)

	fmt.Printf("Wrote %d bytes and %d characters\n", l, len(s))
	return l, nil
}

func main() {
	var c ByteCounter
	f1, _ := os.Open("abc.txt")

	n, _ := io.Copy(&c, f1)

	fmt.Printf("%T %[1]v %v\n", c, n)

}
