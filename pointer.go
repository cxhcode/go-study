package main

import "fmt"

func zeroval(ival int) {
	ival = 1
}

func zeroptr(iptr *int) {
	*iptr = 1
}

func main() {
	i := 10
	zeroval(i)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)
}
