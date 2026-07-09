package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string)
	go uploadFile(ch)

	fileUrl := <-ch
	fmt.Println("File Url", fileUrl)
}

func uploadFile(c chan string) {
	fmt.Println("Uploading file...")
	time.Sleep(3 * time.Second)
	fmt.Println("File upload done!")
	fileUrl := "https://s3.345265.png"
	c <- fileUrl
}
