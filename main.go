//Declare "main" package (package is a way to group functions, and it's made up of all the files in the same directory)
package main

//Import the popular "fmt" package, which contains functions for formatting text, including printing to the console.
import "fmt"

//Implement a "main" function to print a message to the console. A "main" function executes by default when you run the "main" package
func main() {
	//Define the variable "conferenceName" and store the value "Go Conference"
	//Variables in Go must be used in the codebase.
	var conferenceName = "Go Conference"

	//Declare "const", value CAN NOT change
	const conferenceTickets = 50
	var remainingTickets = 50

	//Println(prints the message in a new line)
	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}
