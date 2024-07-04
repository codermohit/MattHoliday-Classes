package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not enough argument")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		fmt.Printf("s : %q", s)
		t := strings.Join(s, new)
		fmt.Printf("%q\n", t)
		fmt.Println(t)
	}
}
