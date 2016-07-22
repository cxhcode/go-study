package main

import "fmt"

func main() {
	var a int = 10
	fmt.Println(a)
	var b int8 = 100
	fmt.Println(b)
	var c = 0
	fmt.Println(c)

	c = 'A'
	fmt.Println(c)

	d := "Jack"
	fmt.Println(d)

	e, f := 1, 2
	fmt.Println(e + f)

	var g, h, _ = 97, "sdfds", 'c'

	fmt.Println(string(g) + h)

	fmt.Println(int64(3e10))
}
