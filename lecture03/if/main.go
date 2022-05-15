package main

import "fmt"

func main() {
	fmt.Println("start of if")
	if 30 > 20 {
		fmt.Println("this code runs because the condition is true")
	}

	fmt.Println("end of if")
	fmt.Println()

	fmt.Println("start of if else")
	if 30 > 20 && 20+5 == 26 {
		fmt.Println("this code runs because the condition is true")
	} else {
		fmt.Println("this code runs because the condition is false")
	}
	fmt.Println("end of if else")
	fmt.Println()

	age := 56

	fmt.Println("start of if else if else")
	if age < 12 {
		fmt.Println("kid")
	} else if age < 18 {
		fmt.Println("not adult")
	} else if age < 40 {
		fmt.Println("adult")
	} else {
		fmt.Println("old")
	}
	fmt.Println("end of if else if else")
	fmt.Println()

}
