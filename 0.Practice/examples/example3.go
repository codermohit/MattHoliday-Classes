package example3

import "fmt"

type Coordinates struct {
	x, y int
}

func (c *Coordinates) Add(c2 *Coordinates) {
	c.x += c2.x
	c.y += c2.y
}

func (c *Coordinates) String() string {
	return fmt.Sprintf("x : %d, y : %d\n", c.x, c.y)
}

func main() {
	c1 := Coordinates{3, 4}
	c2 := Coordinates{7, 6}

	c1.Add(&c2)

	fmt.Println(c1)
}
