package main

import (
	"fmt"
	"io"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

// Implementing sort.Interface
func (s Organs) Len() int {
	return len(s)
}

func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByName struct {
	Organs
}

type ByWeight struct {
	Organs
}

/*
ByName and ByWeight both embed the Organs type, inheriting the Len and Swap methods from Organs. By embedding Organs, both ByName and ByWeight implicitly implement the Len and Swap methods required by the sort.Interface.
*/
func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {
	s := Organs{{"brain", 1340}, {"heart", 240}, {"spleen", 162}, {"pancreas", 131}}
	sort.Sort(ByName{s})
	fmt.Println(s)
	sort.Sort(ByWeight{s})
	fmt.Println(s)

}
