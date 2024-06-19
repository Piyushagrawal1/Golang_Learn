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
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello,", name)

	//using bufio - Print all
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') // jb tk nyi line nahi mil jti tb tk print kr do
	fmt.Println("Hello,", name)

}