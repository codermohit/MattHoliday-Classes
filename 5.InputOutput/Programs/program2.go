package program2

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		data, err := io.ReadAll(file)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		fmt.Fprintf(os.Stdout, "The file has %d bytes\n", len(data))

		file.Close()

	}

}
