package main

import "fmt"

const className string = "programming with go"

var pf = fmt.Printf

func noInput() {
	pf("this is a function with no input or output")
	fmt.Println()
}

// function with input
func hello(name string) {
	pf("hello %s\n", name)
	pf("hello %s\n", name)
	pf("hello %s\n", name)
	fmt.Println(className)
	fmt.Println()
}

// function with output
func add(num1 int, num2 int) int {
	pf("this is a function with output\n")
	return num1 + num2
}

func main() {
	// pf("%+v\n", hello)
	// fmt.Print(hello)
	// hello()
	noInput()
	// fmt.Println("hello world")
	hello("broccoli")
	hello("alan turing")
	hello("euler")

	output := add(2, 3)
	fmt.Print(output)
	fmt.Println()

	formattedString := fmt.Sprintf("this is a formatted string with input: %s\n", "this is input")
	fmt.Print(formattedString)
}
