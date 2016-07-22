package main

import "fmt"

func main() {

	b := make(map[string]int)

	b["a"] = 1
	b["b"] = 2

	fmt.Println(b)
	i, j := b["a"]

	fmt.Println(i, j)

}
