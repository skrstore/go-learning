# Tutorial Point

> Reference : https://www.tutorialspoint.com/go/index.htm

## Overview

- Go language is a programming language initially developed at Google in the year 2007 by 'Robert Griesemer', 'Rob Pike', and 'Ken Thompson'.
- It is a statically-typed language having syntax similar to that of C.
- It provides garbage collection, type safety, dynamic-typing capability, many advanced built-in types such as variable length arrays and key-value maps.
- It also provides a rich standard library.
- The Go programming language was launched in November 2009 and is used in some of the Google's production systems.

## Environment Setup

- VSCode
- Go Installation
- Adding Go Bin path to system path variable

## Hello World Program

```go
// main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

- Command to Run: `go run main.go`

## Basic Syntax

- Go is a case-sensitive programming language.
- We don't need to add semicolon(`;`) after the end of the line.
- Whitespace(blanks, tabs, newline characters, and comments) are ignored by Go Compiler

### Identifiers

- A Go identifier is a name used to identify a variable, function, or any other user-defined item.
- An identifier starts with a letter A to Z or a to z or an underscore \_ followed by zero or more letters, underscores, and digits (0 to 9).

### Keywords

|          |             |        |           |        |
| -------- | ----------- | ------ | --------- | ------ |
| break    | default     | func   | interface | select |
| case     | defer       | Go     | map       | Struct |
| chan     | else        | Goto   | package   | Switch |
| const    | fallthrough | if     | range     | Type   |
| continue | for         | import | return    | Var    |
|          |             |        |           |        |

## DataTypes, Variables, Constants

- Boolean
- Numeric
- String
- Derived
  - Pointer
  - Array
  - Structure
  - Union
  - Function
  - Slice
  - Interface
  - Map
  - Channel

```go
var name = "Admin"
const id = 1001
var course = "GoLang"
var salary int8 = 123

// Static Type declaration
var x float64 = 20.0

// Dynamic type declaration/type inference
// - A dynamic type variable declaration requires the compiler to interpret the type of the variable based on the value passed to it.
y := 42

// Mixed type declaration using type inference
var a, b, c = 3, 4, "foo"
```

### Literals

```go
// Integer Literal
85         /* decimal */
0213       /* octal */
0x4b       /* hexadecimal */
30         /* int */
30u        /* unsigned int */
30l        /* long */
30ul       /* unsigned long */

// Floating-Point Literal
3.14159
314159E-5
12e-3

// String Literals
"hello, dear"
```

### Escaping Characters

- `\n`, `\t`, `\\`

### Type Casting

- Syntax : `type(expression)`

```go
var amount = 10.33
var n = int(amount)
```

## Operators

- Arithmetic - `+, -, *, /, %, ++, --`
- Relational - `==, != , <, >, <=, >=`
- Logical - `&&, ||, !`
- Bitwise
- Assignment
- Miscellaneous - `&, *`
- Operator Precedence

```go
true && false // false
```

## Decision Making Statements

- If statement
- If Else statement
- Nested if statement
- Switch statement
- Select Statement

```go
// If, If Else
if 1 > 2 {
    println("True")
} else {
    println("False")
}

// Switch Statement
const DAY = 1
// Syntax I
switch DAY {
case 1:
    println("Monday")
case 2:
    println("Tuesday")
}

// Syntax 2
switch {
case DAY == 1:
    println("Monday")
case DAY == 2:
    println("Tuesday")
}
// Syntax 3
// Type Switch Statement


// Select Statement
```

## Loops

- For Loop
- Nested Loop
- break Statement
- continue Statement
- goto Statement
- infinite loop

```go
// For , nested loop
package main

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i; j++ {
			print("*")
		}
		println()
	}
}
```

```go
// Break and Continue Statement
package main

func main() {
	for i := 0; i < 5; i++ {
		if i == 1 {
			continue
		}
		if i == 3 {
			break
		}
		println(i)
	}
}
```

```go
// goto statement
package main

func main() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			goto out
		}
		println(i)
	}
out:
	println("Out of the Loop")
}
```

```go
// // infinite Loop
// package main

// func main() {
// 	for true {
// 		println("Infinite")
// 	}
// }
```

## Function

- Parts of a function

  - name
  - parameters
  - return type
  - body

- Function Syntax

```go
func functionName(/* Parameters */) /* Return types*/{
    // Body
}
```

- [Function Examples](../programs/function_examples/readme.md)

## Scope Rules

- Local and global variables are initialized to their default value, which is 0; while pointers are initialized to nil.
- Three places where variable can be declared

  - Inside the function(`local variables`)
  - Outside the function(`global variables`)
  - In definition of function parameters(`formal parameters`)

- [Example Code](../programs/variables_scope/main.go)

## Strings

- Strings are readonly slice of bytes(slices)
- The Go platform provides various libraries to manipulate strings.
  - unicode
  - regexp
  - strings

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// creating a string
	var name string = "Super Admin"
	// name := "Super Admin"
	// var name = "Super Admin"

	fmt.Println("Name : ", name)
	fmt.Printf("Name : %s\n", name)
	// String length
	fmt.Println("Length of string ("+name+") = ", len(name))

	// Hex bytes
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
	fmt.Println()

	// Joining Elements of Array
	nameArr := []string{"Super", "Admin"}
	fmt.Println("Full name - ", strings.Join(nameArr, " "))
}
```

## Arrays

- Array store a fixed-size sequential collection of elements of the same type
- All arrays consist of contiguous memory locations

```go
package main

import "fmt"

func main() {
	// Declaring an array
	// var variable_name [SIZE] variable_type
	var arr [3]int

	// Initializing an array
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	fmt.Println("Arr", arr)

	var balance = [3]float32{123.44, 12.34, 33.44}
	fmt.Println("Balance Array", balance)

	// accessing the elements of array
	for i := 0; i < len(balance); i++ {
		fmt.Printf(" %d --> %f ", i, balance[i])
	}
	fmt.Println()
}
```

### Multi-Dimensional Array

### Passing array to function

```go
package main

import "fmt"

func main() {
	var arr = []int{10, 20, 30}
	var arr1 = [3]int{10, 20, 30}
	printArr(arr)
	fmt.Println("Sum of arr1 : ", sumOfArr(arr1))
}

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf(" %d --> %d ", i, arr[i])
	}
	fmt.Println()
}

// formal parameter with size of array
func sumOfArr(arr [3]int) (sum int) {
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return
}
```

## Pointers

- A pointer is a variable whose value is the address of another variable.
- Array of Pointers
- Pointer to Pointer
- Passing pointers to function

```go
package main

import "fmt"

func main() {
	var name = "Admin"
	var namePointer = &name
	// pointer to pointer
	var ptr = &namePointer

	fmt.Println("Value of name ", name)
	fmt.Println("Value of namePointer", namePointer)
	fmt.Println("Address of namePointer", &namePointer)
	fmt.Println("Value inside namePinter - ", *namePointer)
	fmt.Println("Value of ptr ", ptr)
	fmt.Println("Value inside ptr ", *ptr)
	fmt.Println("Value inside **ptr", **ptr)

	n1, n2 := 12, 34
	fmt.Printf("value before swap n1 : %d, n2 : %d\n", n1, n2)
	swapByAddress(&n1, &n2)
	fmt.Printf("value after swap n1 : %d, n2 : %d\n", n1, n2)
}

// Passing pointers to a function
func swapByAddress(n1 *int, n2 *int) {
	var tmp int = *n1
	*n1 = *n2
	*n2 = tmp
}
```

## Structure

- Structure is another user-defined data type, which allow to combine data items of different kinds.

- [Structure Examples](../programs/structure_examples/main.go)

## Slice

- If a slice is declared with no inputs, then by default, it is initialized as nil. Its length and capacity are zero

- [Slice Examples](../programs/slice_examples/main.go)

## Range

- The range keyword is used in for loop to iterate over items of an array, slice, channel or map.
- For array and slice it returns index
- For map it returns key
- It can return 1 or 2 value(I- Index or key II-value)
- [Range Examples](../programs/range_examples/main.go)

## Map

- It is used to store the date in the format of key and value
- and we can use key to retrieve the value
- Syntax

```go
// var variableName = map[keyDataType] valueDataType
```

- [Map Examples](../programs/map_examples/main.go)

## Recursion

- If a program allows to call a function inside the same function, then it is called a recursive function call.

- [Recursion Examples](../programs/recursion_examples/main.go)

## Interface

- Interfaces represents a set of method signatures. The struct data type implements these interfaces to have method definitions for the method signature of the interfaces.

- [Interface Examples](../programs/interface_examples/main.go)
