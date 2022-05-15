package main

import "fmt"

func main() {
	trueCondition := true
	falseCondition := false
	fmt.Printf("trueCondition: %v\n", trueCondition)
	fmt.Printf("falseCondition: %v\n", falseCondition)
	fmt.Println()
	// not operator !
	fmt.Printf("!falseCondition: %v\n", !falseCondition)
	fmt.Printf("!trueCondition: %v\n", !trueCondition)
	fmt.Println()

	// and operator &&
	fmt.Printf("trueCondition && trueCondition: %v\n", trueCondition && trueCondition)
	fmt.Printf("trueCondition && falseCondition: %v\n", trueCondition && falseCondition)
	fmt.Printf("falseCondition && trueCondition: %v\n", falseCondition && trueCondition)
	fmt.Printf("falseCondition && falseCondition: %v\n", falseCondition && falseCondition)
	fmt.Println()
	// or operator ||
	fmt.Printf("trueCondition || trueCondition: %v\n", trueCondition || trueCondition)
	fmt.Printf("trueCondition || falseCondition: %v\n", trueCondition || falseCondition)
	fmt.Printf("falseCondition || trueCondition: %v\n", falseCondition || trueCondition)
	fmt.Printf("falseCondition || falseCondition: %v\n", falseCondition || falseCondition)
	fmt.Println()

	fmt.Printf("falseCondition || (trueCondition && falseCondition): %v\n", falseCondition || (trueCondition && falseCondition))
	fmt.Printf("falseCondition || (trueCondition && trueCondition): %v\n", falseCondition || (trueCondition && trueCondition))
	fmt.Printf("falseCondition && trueCondition && falseCondition: %v\n", falseCondition && trueCondition && falseCondition)

}
