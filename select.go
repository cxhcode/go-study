package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 4)
		c1 <- "ping"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "pong"
	}()

	//for i := 0; i < 2; i++ {
	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(time.Second * 3):
		fmt.Println("time out")
	}

}


