package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanWords)

	words := make(map[string]int)

	type kv struct {
		key string
		val int
	}

	var ss []kv

	for scanner.Scan() {
		words[scanner.Text()]++
	}

	/*
	   Code ahead takes all the key value pairs from the 'words' map
	   puts them in the 'ss' slice and then sorts the slice on the basis
	   of the number of times a key has occured.
	*/
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, kv := range ss[:4] {
		fmt.Println(kv.key, " : ", kv.val)
	}

}
