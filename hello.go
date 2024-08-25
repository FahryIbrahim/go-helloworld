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
