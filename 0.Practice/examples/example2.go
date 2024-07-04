package example2

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// uncomment below line to read word by word
	//scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Fprintln(os.Stdout, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("%v", err)
	}
}
