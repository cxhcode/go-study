package main

import "fmt"

func main() {
	a := 10
	if a > 5 {
		fmt.Println(a)
	} else {
		fmt.Println("")
	}

	if b := 4; b < 2 {
		fmt.Println(b)
	}
}
