package main

import "fmt"

// Global variable
var name = "Super Admin"

func main() {
	fmt.Println("Variable Scope")

	// Local variable
	name := "Admin"

	fmt.Println("Name : ", name)
	// actual parameter the value
	run(name)
}

// Function with formal parameter(treated as local variable)
func run(user string) {
	fmt.Println("Run by : ", user)

	fmt.Println("From run() name =", name)
}
