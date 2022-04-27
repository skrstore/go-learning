package main

import "fmt"

func main() {
	var arr = []int{10, 20}

	// slice
	arr2 := make([]int, 0, 6)
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", arr2)
	fmt.Println(arr2)
	fmt.Println(arr[1:2])

	printArrInfo(arr)
	printArrInfo(arr2)

	// Nil slice
	var numbers []int

	fmt.Println("Nil Slice")
	printArrInfo(numbers)

	// Subslicing
	alphabets := []string{"A", "B", "C", "D", "E"}
	fmt.Println("alphabets[1:] ", alphabets[1:])
	fmt.Println("alphabets[1:3] ", alphabets[1:3])
	fmt.Println("alphabets[1:3] ", alphabets[:])

	fmt.Println("------- append() and copy()")

	var numbers1 []int
	printArrInfo(numbers1)
	numbers1 = append(numbers1, 12)
	printArrInfo(numbers1)
	numbers1 = append(numbers1, 12)
	printArrInfo(numbers1)

	var numbers2 = []int{11, 22, 33}
	copy(numbers1, numbers2) // replace the old elements with new
	fmt.Print("Numbers 1 ")
	printArrInfo(numbers1)
	fmt.Print("Numbers 2 ")
	printArrInfo(numbers2)
}

func printArrInfo(arr []int) {
	// Length and Cap(capacity) of slice or array
	fmt.Printf("len = %d cap = %d elem = %v\n", len(arr), cap(arr), arr)
}
