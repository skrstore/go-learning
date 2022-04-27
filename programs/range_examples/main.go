package main

import "fmt"

func main() {
	numbers := []int{11, 22, 33, 44, 55}

	for index, value := range numbers {
		fmt.Println(index, numbers[index], value)
	}

	courses := map[string]string{"frontend": "ReactJS", "backend": "NodeJS"}

	for course := range courses {
		fmt.Printf("key : %s : value : %s\n", course, courses[course])
	}

	for key, value := range courses {
		fmt.Printf("key : %s : value : %s\n", key, value)
	}
}
