//Declare "main" package (package is a way to group functions, and it's made up of all the files in the same directory)
package main

//Import the popular "fmt" package, which contains functions for formatting text, including printing to the console.
import "fmt"

//Implement a "main" function to print a message to the console. A "main" function executes by default when you run the "main" package
func main() {
	//Println(prints the message in a new line)
	fmt.Println("Welcome to our booking application")
	fmt.Println("Get your tickets here to attend")
	//Define the variable "conferenceName" and store the constant value "Go Conference"
	var conferenceName = "Go Conference"
	//Variables in Go must be used in the codebase. The line below uses the variable.
	fmt.Println(conferenceName)
}
