package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	var start = time.Now()

	// wg.Add(1)
	// go uploadFile()
	wg.Go(uploadFile)
	// wg.Add(1)
	// go saveToDB()
	wg.Go(saveToDB)
	// wg.Add(1)
	// go sendEmail()
	wg.Go(sendEmail)

	// time.Sleep(3 * time.Second)
	wg.Wait() // waiting .... until counter is 0
	fmt.Println("All tasks completed")
	fmt.Println("time taken", time.Since(start))
}

func uploadFile() {
	// defer wg.Done()
	fmt.Println("Uploading file...")
	time.Sleep(3 * time.Second)
	fmt.Println("File upload done!")
	// wg.Add(-1)
	// wg.Done()

}

func saveToDB() {
	// defer wg.Done()
	fmt.Println("saving to db...")
	time.Sleep(1 * time.Second)
	fmt.Println("Saved to DB!")
	// wg.Add(-1)
	// wg.Done()

}

func sendEmail() {
	// defer wg.Done()
	fmt.Println("Sending email...")
	time.Sleep(2 * time.Second)
	fmt.Println("email sent!")
	// wg.Add(-1)
	// wg.Done()

}
