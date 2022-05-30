//Declare "main" package (package is a way to group functions, and it's made up of all the files in the same directory)
package main

//Import the popular "fmt" package, which contains functions for formatting text, including printing to the console.
import "fmt"

//Implement a "main" function to print a message to the console. A "main" function executes by default when you run the "main" package
func main() {
	//Define the variable "conferenceName" and store the value "Go Conference"
	//Variables in Go must be used in the codebase.
	//Go ASSUMES the data type of the value if no data type is assigned.

	//Alternative syntax for creating a variable | Constants and declaring data types are not possible
	conferenceName := "Go Conference"

	//var conferenceName = "Go Conference"

	//Declare "const", value CAN NOT change
	const conferenceTickets = 50
	var remainingTickets = 50

	//Printf(Prints formatted data, some text with a variable %s, myVariable)
	//Added \n at end of sentence to create a new line
	// fmt.Printf("conferenceTickets is %T, remaining tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	//Println(prints the message in a new line)
	fmt.Println("Get your tickets here to attend")

	//ask user for their name

	//in Go variables without a value must be assigned a data type
	//String data types generally for text based
	//Int(Integer) data types generally for whole numbers
	var firstName string
	var lastName string
	var email string
	var userTickets int
	//Scan interprets user input
	//Pointer is a variable that points to the memory address of another variable
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
}
