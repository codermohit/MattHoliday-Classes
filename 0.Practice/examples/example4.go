package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var output io.WriteCloser
	var err error

	output, err = os.Create("abc.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open file : %s", err)
	}

	defer output.Close()

	scanner := bufio.NewScanner(os.Stdout)

	for scanner.Scan() {
		str := scanner.Text()
		fmt.Fprintf(output, "%s\n", str)
	}

	fmt.Fprintf(os.Stdout, "%T\n", output)
}
