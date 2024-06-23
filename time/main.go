package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Time in GO")

	currTime := time.Now()
	fmt.Println("Time is: ", currTime)

	fmt.Println("Date is: ", currTime.Format("01-02-2006"))  //this is fixed format

	fmt.Println("Date is: ", currTime.Format("01-02-2006 03:04:05 PM"))    //this is fixed format
}
