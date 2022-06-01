//Declare "main" package (package is a way to group functions, and it's made up of all the files in the same directory)
package main

//Import the popular "fmt" package, which contains functions for formatting text, including printing to the console.
//Imported strings package to split the string with white space as separator, returns a slice with the split elements
import (
	"fmt"
	"strings"
)

//Implement a "main" function to print a message to the console. A "main" function executes by default when you run the "main" package
func main() {
	//Define the variable "conferenceName" and store the value "Go Conference"
	//Variables in Go must be used in the codebase.
	//Go ASSUMES the data type of the value if no data type is assigned.

	//Alternative syntax for creating a variable | Constants and declaring data types are not possible
	conferenceName := "Go Conference"

	//var conferenceName = "Go Conference"

	//Declare "const", value CAN NOT change
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	//Arrays are data structures to store collection of elements in a single variable
	//Arrays have a fixed size(how many elements the array can hold)
	//Only can use the stored data type
	//var bookings = [50]string{}
	//Alternate syntax | var bookings [50]string
	var bookings []string
	//Alternate syntax for slice | bookings := []string{}

	//Printf(Prints formatted data, some text with a variable %s, myVariable)
	//Added \n at end of sentence to create a new line
	// fmt.Printf("conferenceTickets is %T, remaining tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	//Println(prints the message in a new line)
	fmt.Println("Get your tickets here to attend")

	//Problems with Arrays
	//Maybe 1 user books all tickets
	//Having an array with size 50, with only 1 element inside is not efficient
	//What if we don't know the size when creating it?

	//ask user for their name
	//Created for loop to loop the booking application
	//Added condition below to run loop as long as it is met
	for remainingTickets > 0 && len(bookings) < 50 {
		//in Go variables without a value must be assigned a data type
		//String data types generally for text based
		//Int(Integer) data types generally for whole numbers
		var firstName string
		var lastName string
		var email string
		var userTickets uint
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

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			//fmt.Printf("The whole slice: %v\n", bookings)
			//fmt.Printf("The first value: %v\n", bookings[0])
			//fmt.Printf("Slice type: %T\n", bookings)
			//fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			//Added "for each" loop nested inside for loop
			//For each full entry, extract only first names from the bookings
			//Created Slice with empty values named firstNames
			//_ (underscore) Blank idenitifier | ignores a variable you don't want to use
			//Range iterates over elements for different data structures (Not only arrays and slices)
			//For arrays and slices, range provides the index and value for each element
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			//var noTicketsRemaining bool = remainingTickets == 0
			//noTicketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				//End the program
				fmt.Println("Our conference is booked. Please come again next time.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			//continue causes loop to skip the remainder of its body
			//immediately retesting its condition
			//continue
			//else if userTickets == remainingTickets
			//You can have unlimited "else if" statments
			//Only 1 "if" statement and 1 "else" statement
		}
	}
}
