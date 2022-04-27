package main

import "fmt"

type studentDetail struct {
	id     int
	name   string
	course string
}

type courseDetail struct {
	// First letter capital
	Name     string
	Duration int
}

func main() {
	var s1 studentDetail = studentDetail{101, "Rahul", ".Net"}
	s1.course = ".NetCore"

	fmt.Println("S1 : ", s1)
	printStudentDetail(s1)
	// Structure Pointer
	printStudentDetailWithUpdate(&s1)
	fmt.Println("After Change")
	printStudentDetail(s1)
	fmt.Println("Update 2")
	printStudentDetailUpdate2(s1)
	fmt.Println("After local Change")
	printStudentDetail(s1)

	fmt.Println("Course Detail Struct")
	course := courseDetail{Name: "Full Stack Web Deeloper", Duration: 1}

	fmt.Println("Course ", course)
}

// passing struct to function argument
func printStudentDetail(student studentDetail) {
	fmt.Println("Studnet ID     : ", student.id)
	fmt.Println("Student Name   : ", student.name)
	fmt.Println("Student Course : ", student.course)
}

// passing struct to function argument
func printStudentDetailWithUpdate(student *studentDetail) {
	student.id = 1001
	fmt.Println("Studnet ID     : ", student.id)
	fmt.Println("Student Name   : ", student.name)
	fmt.Println("Student Course : ", student.course)
}

// passing struct to function argument
func printStudentDetailUpdate2(student studentDetail) {
	student.id = 10003
	fmt.Println("Studnet ID     : ", student.id)
	fmt.Println("Student Name   : ", student.name)
	fmt.Println("Student Course : ", student.course)
}
