package main

import "fmt"

func main() {
	messages := make(chan string)

	//go func() {
	//	messages <- "ping"
	//	messages <- "Jack Chan"
	//}()
	go write(messages)

	msg := <-messages
	//msg1 := <-messages
	fmt.Println(msg)
	//fmt.Println(msg1)
}

func write(messages chan string) {
	messages <- "Jack Chan"
}