package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Day()
	// day := 23
	// fmt.Println(date)
	switch day % 4 {
	case 0:
		fmt.Println("we had a class yesterday")
	case 1:
		fmt.Println("no class today")
	case 2:
		fmt.Println("we will have a class tomorrow")
	case 3:
		fmt.Println("we have a class today")
	}
}
