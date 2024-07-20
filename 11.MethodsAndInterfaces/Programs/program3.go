package main

import "fmt"

type Pair struct {
	Path string
	Hash string
}

type PairWithLength struct {
	Pair
	Length int
}

type Fizgig struct {
	*PairWithLength
	Broken bool
}

func (p Pair) String() string {
	return fmt.Sprintf("Path is : %v, Hash : %v\n", p.Path, p.Hash)
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Path is : %v, Hash : %v, Length : %d\n", p.Path, p.Hash, p.Length)
}

func main() {
	p := Pair{"/usr/bin", "0x4f3f"}
	pwl := PairWithLength{p, 133}

	fzg := Fizgig{
		&PairWithLength{Pair{"/usr/lib/", "0x44dead"}, 12},
		true,
	}

	fmt.Println(p, pwl)
	//Promotion of method
	fmt.Printf("Fizgig : %v", fzg)
}
