package main

import "fmt"

func goruntine(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str)
	}
}

func main() {
	goruntine("Chan")

	go goruntine("Jack")

	go func(str string) {
		fmt.Println();
	}("ChenXinHua")

	var input int

	fmt.Scanln(&input)
	fmt.Println("done")
}
