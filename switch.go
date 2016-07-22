package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println("variable is " + string(i))
	}

	var time1 = time.Now().Weekday()
	fmt.Println(time1)
	switch  time1{
	case time.Saturday, time.Sunday:
		fmt.Println(time.Now().Weekday())
	default:
		fmt.Println("默认值输出")
	}
}
