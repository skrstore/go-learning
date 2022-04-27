# Function Examples

## Function with parameters and return value

```go
package main

func main() {

	println("Is 12 Even : ", isEven(12))
}

func isEven(number int) bool {
	return (number%2 == 0)
}
```

## Function with multiple parameters

```go
package main

import "fmt"

func main() {
	var n1, n2 = 23, 45

	fmt.Printf("Max of %d and %d is %d\n", n1, n2, max(n1, n2))
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
```

## Function as value

```go
package main

import "fmt"

func main() {

	add := func(n1 int, n2 int) int {
		return n1 + n2
	}

	fmt.Println("1 + 2 =", add(1, 2))
}
```

## Function with multiple return value

```go
package main

import "fmt"

func main() {
	var n1, n2 = 23, 45

	fmt.Printf("Value before swap %d and %d\n", n1, n2)
	n1, n2 = swap(n1, n2)
	fmt.Printf("Value after swap %d and %d\n", n1, n2)

}

func swap(n1 int, n2 int) (int, int) {
	return n2, n1
}
```

## Function with named return values

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fileFullName := "readme.md"
	name, ext := splitNameExt(fileFullName)
	fmt.Printf("Filename - %s\nExtension - %s\n", name, ext)
}

func splitNameExt(fileFullName string) (fileName, extension string) {

	splitName := strings.Split(fileFullName, ".")
	fileName = splitName[0]
	extension = splitName[1]
	return
}
```

## [Function as method](./lib/function_as_method.go)
