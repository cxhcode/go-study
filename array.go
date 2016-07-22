package main

import "fmt"

func main() {

	i := [10]int{1, 2, 34, 4, 4, 0, 20}
	fmt.Println(i)

	var a = new([5]int)
	test(a)
	fmt.Println(a)

	var c = new(int)

	var b = [10]*int{c}
	b[1] = *2
	fmt.Println(b)
}

func test(a *[5]int) {
	a[3] = 23
}