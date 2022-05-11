package main

import "fmt"

func main() {
	// var students string = "student1, student2, student3, ..."

	// arrayName = [numberOfItems] dataType {items ...}
	var students = [3]string{"Naruto", "Sasuke"}
	// [Naruto Sasuke Sakura]

	fmt.Print("students: ", students, "\n")

	fmt.Print("students[0]: ", students[0], "\n")
	fmt.Print("students[1]: ", students[1], "\n")
	fmt.Print("students[2]: ", students[2], "\n")
	fmt.Print("len(students): ", len(students), "\n")
	lengthOfStudent := len(students)
	fmt.Println(lengthOfStudent)
	// fmt.Print("students[2]: ", students[2])
}
