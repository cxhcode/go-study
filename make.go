package main

import "fmt"

func main() {

	ints := make([]int, 2, 8)
	ints[0] = 100
	fmt.Println(ints)
	fmt.Println(len(ints), cap(ints))

	ints = append(ints, 1224)
	fmt.Println(ints)

	ints2 := make([]int, len(ints))

	copy(ints2, ints)

	fmt.Println(ints2)
}
