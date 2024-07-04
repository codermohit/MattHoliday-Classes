package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {
	var emps = map[string]*Employee{}

	emps["Mohit"] = &Employee{
		"Mohit", 1, nil,
		time.Now(),
	}

	emps["Jonas"] = &Employee{
		Name:   "Jonas",
		Number: 2,
		Boss:   emps["Mohit"],
		Hired:  time.Now(),
	}

	fmt.Printf("%T %[1]p %[2]p\n", emps["Jonas"], emps["Jonas"].Boss)
	fmt.Printf("%T %[1]p\n", emps["Mohit"])
}
