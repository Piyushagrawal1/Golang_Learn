package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello world")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Hello again")
}

func sayHi() {
	fmt.Println("Say Hi Bhai")
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	fmt.Println("Learning Goroutine...")
	go sayHello()
	go sayHi()

	//wait for a moment to allow the goroutine to finish

	time.Sleep(1000 * time.Millisecond)

}
