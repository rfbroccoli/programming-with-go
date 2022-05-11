package main

import "fmt"

func main() {
	num1 := 20
	num2 := 20
	fmt.Printf("num1:%d\n", num1)
	fmt.Printf("num2:%d\n", num2)

	fmt.Print("num1>num2:", num1 > num2, "\n")
	fmt.Print("num1<num2:", num1 < num2, "\n")
	fmt.Print("num1==num2:", num1 == num2, "\n")
	fmt.Print("num1!=num2:", num1 != num2, "\n")
	fmt.Print("num1<=num2:", num1 <= num2, "\n")
	fmt.Print("num1>=num2:", num1 >= num2, "\n")
}
