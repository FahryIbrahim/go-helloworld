# Hello World in Go

This is a simple "Hello, World!" program written in Go. It demonstrates the basic structure of a Go program, including package import, function definition, and output to the console. Additionally, it reads and displays the current user's environment variable from the operating system.

## Program Overview

The program performs the following tasks:

1. **Prints "Hello, World!"**: A basic string is printed to the console.
2. **Prints a personalized message**: The program reads the `$USER` environment variable from the operating system and prints a message including the user's name.

## Code Explanation

```go
// Hello World in Go by Vivek Gite
package main

// Import OS and fmt packages
import (
	"fmt"
	"os"
)

// Main function
func main() {
	fmt.Println("Hello, World!") // Print simple text on screen
	fmt.Println(os.Getenv("USER"), ", Lets be friends!") // Read Linux $USER environment variable
}
