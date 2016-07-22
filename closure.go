package main

import "fmt"

func main() {

	out1 := out()
	println(out1())
	println(out1())
	println(out1())
	out2 := out()
	println(out2())

	var fc func() int
	fc = func() int{
		return 2
	}

	fmt.Println(fc())
}

func out() func() int {
	num := 0
	return func() int {
		num = num + 1
		return num
	}
}
