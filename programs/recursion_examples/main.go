package main

import "fmt"

func main() {
	n := 6
	fmt.Printf("Factorial of %d is %d\n", n, factorial(n))
	fmt.Printf("%d th fibonacci is %d\n", n, fibonacci(n))

	fmt.Println("Fibonacci series")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
