package main

import (
	"fmt"
	"time"
)

func main() {

	var start = time.Now()
	go uploadFile()
	go saveToDB()
	go sendEmail()

	time.Sleep(3 * time.Second)
	fmt.Println("All tasks completed")
	fmt.Println("time taken", time.Since(start))
}

func uploadFile() {
	fmt.Println("Uploading file...")
	time.Sleep(3 * time.Second)
	fmt.Println("File upload done!")
}

func saveToDB() {
	fmt.Println("saving to db...")
	time.Sleep(1 * time.Second)
	fmt.Println("Saved to DB!")
}

func sendEmail() {
	fmt.Println("Sending email...")
	time.Sleep(2 * time.Second)
	fmt.Println("email sent!")
}
