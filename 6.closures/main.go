package main

import "fmt"

func do(d func()) {
	d()
}

func main() {
	var ss []func()

	/* Closures
	   for loop instantiates i , every iteration
	   if you initialize i before for loop then the location of i remains same
	*/
	//var i int
	for i := 0; i < 5; i++ {
		i2 := i
		v := func() {
			fmt.Printf("%d @ %p\n", i2, &i2)
		}

		ss = append(ss, v)
	}

	for i := 0; i < 5; i++ {
		do(ss[i])
	}

}
