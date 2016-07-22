package main

import "fmt"

func main() {
	a, b := 1, 3
	c, d := swap(a, b)
	fmt.Println(c, d)

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 78))
}

func swap(a int, b int) (int, int) {
	return b, a
}

func sum(num ...int) int {
	var sum int
	for idx, value := range num {
		println("idx=", idx, "value=", value)
		sum += value
		println(sum)
	}
	return sum
}