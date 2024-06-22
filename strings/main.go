package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Split String
	data := "apple, banana, mango, orange"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	//2. Count Occurence

	str := "two one, three four two twp five two"
	count := strings.Count(str, "two")

	fmt.Println(count)

	//3. trim

	str2 := "  Hello, World!  "
	trimed := strings.TrimSpace(str2)
	fmt.Println(trimed)

	//upar vale me to variable bna liya phle fir use print kr liya
	//niche vale me direct print kr diya
	fmt.Println(strings.Trim(str2, " "))   //1st way to trim
	fmt.Println(strings.TrimSpace(str2))   //2nd way to trim
	

	str3 := "piyush"
	str4 := "agrawal"
	// result := strings.Join([]string{str3, str4}, " ")  //1st way to do this
	// fmt.Println(result)

	result := []string{str3, str4}    //2nd way to do this
	fmt.Println(strings.Join(result, " "))
}
