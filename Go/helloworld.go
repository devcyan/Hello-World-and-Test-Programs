// This program should print "Hello World!"

// Packages group functions together similarly to how namespaces group classes in C#
package main

// Importing fmt is required to print output to the terminal.
import "fmt"

// To execute code, a function has to be defined using the func keyword.
// fmt.Println() is the imported function to print output to the terminal.
func main() {
	fmt.Println("Hello World!")
}

// The main function has a special property.
// Main functions act as a starting point call upon other functions defined elsewhere in your code.
// They must also be stored within the main package.
// So, it is important to structure your code in a way that utilizes the main function.
