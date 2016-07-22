package main

import "fmt"

type people struct {
	name string
	age  int
}

func main() {
	people := people{name:"Jack", age:20}
	fmt.Println(people)
	people.age = 21
	fmt.Println(people)
	(&people).age = 25
	fmt.Println(people)

	fmt.Println(&people)

	i := 10
	fmt.Println(&i)
}
