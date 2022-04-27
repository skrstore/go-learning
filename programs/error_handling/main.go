package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("negative number not alloted")
	}
	return int(math.Sqrt(float64(n))), nil
}

func main() {
	n := 16

	sqrtValue, err := sqrt(n)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sqrt of %d = %d\n", n, sqrtValue)
	}
}
