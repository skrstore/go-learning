package main

import "fmt"

func main() {
	// creating a map
	var employeeInfo = map[string]string{"name": "Admin", "department": "IT", "email": "admin@gmail.com"}

	fmt.Println(employeeInfo)

	name, ok := employeeInfo["name"]
	department := employeeInfo["department"]
	fmt.Println("Employee Name ", name, ok)
	fmt.Println("Employee Department ", department)

	fmt.Println("before Delete")
	for key, value := range employeeInfo {
		fmt.Printf("[%s] --> %s\n", key, value)
	}

	delete(employeeInfo, "email")
	fmt.Println("After Delete")
	for key, value := range employeeInfo {
		fmt.Printf("[%s] --> %s\n", key, value)
	}
}
