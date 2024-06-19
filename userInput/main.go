package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey, What's your name?")
	//taking user input
	//it print only before spaces
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello,", name)

	

}