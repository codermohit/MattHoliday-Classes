package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var tlc, twc, tcc int
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		var lc, wc, cc int

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}

		tlc += lc
		twc += wc
		tcc += cc

		fmt.Fprintf(os.Stdout, "%7d %7d %7d %s\n", lc, wc, cc, fname)

		file.Close()

	}

	fmt.Fprintf(os.Stdout, "%7d %7d %7d total\n", tlc, twc, tcc)
}
