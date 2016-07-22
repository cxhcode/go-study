package main

import "fmt"

func main() {

	ints := []int{2, 34, 3}
	for idx, item := range ints {
		fmt.Println("idx = ", idx, "item = ", item)
	}

	for idx, item := range "AaHello" {
		fmt.Println("idx = ", idx, "item = ", item)
	}

}
