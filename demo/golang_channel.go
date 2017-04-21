package main

import "fmt"

func GolangChannel() {
	msgChan := make(chan string)
	go func() {
		msgChan <- "Hello, golang channel!"
	}()

	msg := <-msgChan
	fmt.Println(msg)
}
