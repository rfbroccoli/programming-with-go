package main

import "fmt"

func main() {
	// var students string = "student1, student2, student3, ..."

	// sliceName = [] dataType {items ...}
	var students = []string{"Naruto", "Sasuke", "Sakura", "Shikamaru"}
	// [Naruto Sasuke Sakura]

	fmt.Print("students: ", students, "\n")
	fmt.Print("students[0]: ", students[0], "\n")
	fmt.Print("students[1]: ", students[1], "\n")
	fmt.Print("students[2]: ", students[2], "\n")
	fmt.Print("students[3]: ", students[3], "\n")
	fmt.Print("len(students): ", len(students), "\n")

	// fmt.Print("students[2]: ", students[2])
}
