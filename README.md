# Learn Go

## Important Commands

```sh
# Running a file
go run hello.go

# Running a Module
go run .

# Building the Binary
go build .

# To create a module
go mod init MODULE_NAME
```

## Reference Material

- https://golang.org/doc/
- [Go Basics - Steve Hook - YouTube - Playlist](https://www.youtube.com/playlist?list=PLsc-VaxfZl4cxYWszvnmmWY9H5Jim79R2)
- https://github.com/inancgumus/learngo
- https://github.com/mmcgrana/gobyexample
- https://github.com/kkdai/project52
- https://www.youtube.com/c/GolangDojo/playlists
- https://tour.golang.org/welcome/1
- https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go
- [Tutorial Point - Notes](./tutorials/tutorial_point.md)

> Testing

- Testing APIs - https://tutorialedge.net/golang/advanced-go-testing-tutorial/
- https://gobyexample.com/testing
- https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package

> Logging

- https://stackoverflow.com/questions/16895651/how-to-implement-level-based-logging-in-golang
- https://www.honeybadger.io/blog/golang-logging/
- https://github.com/sirupsen/logrus
- https://github.com/uber-go/zap

> Go Projects to Explore

- https://github.com/brandur/sorg
- https://github.com/stretchr/testify
- https://github.com/gin-gonic/gin

## Notes

### Program Structure

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
    println("Hello, World")
}
```

### Taking input from User

```go
package main

import "fmt"

func main() {
	print("Enter your name : ")
	var name string
	fmt.Scanln(&name)

	println("Hello, ", name)
}
```

### Packages

- Packages Types

  - Executable packages(main) - these contain the main function(package)
  - Non Executable package - these does not contain main function

- package name should be same as Directory name
- package name should be short and without underscore and camel case they should be small case letter

### If statement with initialization

> Reference - https://golangdocs.com/the-if-else-expression-in-golang

```go
package main

import (
    "fmt"
)

func v() int {
    return 42
}

func main() {
    if a := v(); a == 42 {
        fmt.Println("It's 42")        // It's 42
    } else {
        fmt.Println("It's not")
    }
}
```

### Variadic Functions - Variable length argument function

```golang
package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	numbers := []int{1, 2, 3, 4}
	fmt.Println(sum(numbers...))
}
// Variadic Function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
```

### To Build binary for different os

> Reference - https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures

- Building from windows for linux
  - `GOOS=linux go build main.go`
- Building from linux for Windows
  - `GOOS=windows go build main.go`
  - `GOOS=windows go build`

---

## YouTube Videos Notes

### What can you build in Golang? - Golang Dojo

> Reference : https://youtu.be/4fjNO9CuqVs

Notes

- Devops Tooling
- Backend Web Services(using native libraries)
- System programming
- Blockchain Client
- WebAssembly
- Web Site Development(template library)
- Note To develop
  - Trading algorithm
  - Game Development

---

## Todo

- How to call API using go.
- How to connect to DB(MySQL, MongoDB, SQLite)
- How to manage third party packages(installing globally, in vendor folder)
- How to send Mail
- How to read environment variables
- how to create a simple server
- how to CRUD on files, dirs
- how to parse URL, path
