package main

import (
	"fmt"
	"time"
)

func main() {

	var ch = make(chan string, 3)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "file uploaded!"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "file url saved!"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "email sent!"
	}()

	for range 3 {
		data := <-ch
		fmt.Println(data)
	}
}
