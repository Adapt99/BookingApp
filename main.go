//Declare "main" package (package is a way to group functions, and it's made up of all the files in the same directory)
package main

//Import the popular "fmt" package, which contains functions for formatting text, including printing to the console.
//Imported strings package to split the string with white space as separator, returns a slice with the split elements
import (
	//"BookApp1-techworldnana/helper"
	"fmt"
	"strings"
)

//Package level variables
//Defined at the top outside all functions
//They can be accessed inside any of the functions
//And in all files, which are in the same package

//Best Practice: define variable as "local" as possible | Create the variable where you need it
//Local Variables | Defined inside a function or a block | They can be accessed only sinde that function or block of code
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {
	//Define the variable "conferenceName" and store the value "Go Conference"
	//Variables in Go must be used in the codebase.
	//Go ASSUMES the data type of the value if no data type is assigned.

	//Alternative syntax for creating a variable | Constants and declaring data types are not possible
	//conferenceName := "Go Conference"

	//var conferenceName = "Go Conference"

	//Declare "const", value CAN NOT change
	//const conferenceTickets int = 50
	//var remainingTickets uint = 50

	//Arrays are data structures to store collection of elements in a single variable
	//Arrays have a fixed size(how many elements the array can hold)
	//Only can use the stored data type
	//var bookings = [50]string{}
	//Alternate syntax | var bookings [50]string
	//var bookings []string
	//Alternate syntax for slice | bookings := []string{}

	greetUsers()

	//Printf(Prints formatted data, some text with a variable %s, myVariable)
	//Added \n at end of sentence to create a new line
	//fmt.Printf("conferenceTickets is %T, remaining tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	//fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	//fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	//Println(prints the message in a new line)
	//fmt.Println("Get your tickets here to attend")

	//Problems with Arrays
	//Maybe 1 user books all tickets
	//Having an array with size 50, with only 1 element inside is not efficient
	//What if we don't know the size when creating it?

	//ask user for their name
	//Created for loop to loop the booking application
	//Added condition below to run loop as long as it is met
	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()
		//in Go variables without a value must be assigned a data type
		//String data types generally for text based
		//Int(Integer) data types generally for whole numbers
		//var firstName string
		//var lastName string
		//var email string
		//var userTickets uint
		//Scan interprets user input
		//Pointer is a variable that points to the memory address of another variable
		//fmt.Println("Enter your first name: ")
		//fmt.Scan(&firstName)

		//fmt.Println("Enter your last name: ")
		//fmt.Scan(&lastName)

		//fmt.Println("Enter your email address: ")
		//fmt.Scan(&email)

		//fmt.Println("Enter number of tickets: ")
		//fmt.Scan(&userTickets)

		//Validate user input of first and last name | Length of names entered must be greater or equal to 2
		//var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		//Validate user input of email | Email must contain @ symbol
		//isValidEmail := strings.Contains(email, "@")
		//Validate user input of ticket number | Number must be positive and greater than 0 AND less than or equal to the amount of remaining tickets
		//isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
		//isValidCity := city == "Singapore" || city == "London"
		//!isValidCity
		//isInvalidCity := city != "Singapore" || city != "London"

		//Modified if statement to include user input validation
		//Application will not run if criteria isn't met
		if isValidName && isValidEmail && isValidTicketNumber {
			//remainingTickets = remainingTickets - userTickets
			//bookings = append(bookings, firstName+" "+lastName)

			bookTicket(userTickets, firstName, lastName, email)
			//fmt.Printf("The whole slice: %v\n", bookings)
			//fmt.Printf("The first value: %v\n", bookings[0])
			//fmt.Printf("Slice type: %T\n", bookings)
			//fmt.Printf("Slice length: %v\n", len(bookings))

			//fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			//fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			//Added "for each" loop nested inside for loop
			//For each full entry, extract only first names from the bookings
			//Created Slice with empty values named firstNames
			//_ (underscore) Blank idenitifier | ignores a variable you don't want to use
			//Range iterates over elements for different data structures (Not only arrays and slices)
			//For arrays and slices, range provides the index and value for each element
			//firstNames := []string{}
			//for _, booking := range bookings {
			//	var names = strings.Fields(booking)
			//	firstNames = append(firstNames, names[0])
			//}
			//fmt.Printf("The first names of bookings are: %v\n", firstNames)

			firstNames := getFirstNames()
			fmt.Printf("the first names of bookings are: %v\n", firstNames)
			//var noTicketsRemaining bool = remainingTickets == 0
			//noTicketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				//End the program
				fmt.Println("Our conference is booked. Please come again next time.")
				break
			}
		} else {
			//Added Error messages to display proper feedback
			if !isValidName {
				fmt.Println("First name or last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address entered doesn't contain @ symbol")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets entered is invalid")
			}
			//fmt.Println("Your input data is invalid, please try again.")
			//fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			//continue causes loop to skip the remainder of its body
			//immediately retesting its condition
			//continue
			//else if userTickets == remainingTickets
			//You can have unlimited "else if" statments
			//Only 1 "if" statement and 1 "else" statement
		}
	}
	//Switch statements
	//Based on selected city, different code will be executed
	//city := "London"
	//switch city {
	//case "New York":
	//execute code
	//case "Singapore", "Hong Kong":
	//execute code for Singapore & Hong Kong
	//case "London":
	//execute code
	//default:
	//fmt.Print("No valid city selected")
	//}
}

//Function is only executed, when "called"
//You can call a function infinitely
//Functions are used to reduce code duplication

//Paremeters | Information can be passed into functions as parameters
//Parameters are also called arguments
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	//A function can "return" data as a result
	//So a function can take an input and return an output
	//In Go you have to define the input and output parameters including its type explicitly
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	//A Go function can return multiple values
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	//fmt.Printf("The whole slice: %v\n", bookings)
	//fmt.Printf("The first value: %v\n", bookings[0])
	//fmt.Printf("Slice type: %T\n", bookings)
	//fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
