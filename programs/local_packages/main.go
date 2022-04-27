package main

import (
	// built-in
	"fmt"

	// local
	"local_packages/lib"
)

func main() {
	fmt.Println("Local packages")

	result := lib.Add(1, 3)

	fmt.Println("1 + 3 = ", result)
}
