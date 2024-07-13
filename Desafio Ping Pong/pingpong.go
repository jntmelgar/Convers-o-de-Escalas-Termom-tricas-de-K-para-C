package main

import (
	"fmt"
	"time"
)

func ping(p chan string, c chan string) {
	for i := 0; ; i++ {
		p <- "ping"
		c <- "pong"
	}
}
func imprimir(p chan string, c chan string) {
	for {
		msg1 := <-p
		msg2 := <-c
		fmt.Println(msg1)
		time.Sleep(500 * time.Millisecond)
		fmt.Println(msg2)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	p := make(chan string)
	c := make(chan string)

	go ping(p, c)
	go imprimir(p, c)
	var escreva string
	fmt.Scanln(&escreva)
}
