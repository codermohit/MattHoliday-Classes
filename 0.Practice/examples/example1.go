package example1

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var n int

	for {

		var val float64
		if _, err := fmt.Fscanln(os.Stdin, &val); err != nil {
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "No values")
		os.Exit(-1)
	}

	fmt.Fprintf(os.Stdout, "Average value is : %f\n", sum/float64(n))
}
