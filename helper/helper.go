package helper

//Go programs are organized into packages
//A package is a collection of Go files

import "strings"

//Exporting a variable | Makes it available for all packages in the application
//Capitalize the first letter

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	//A Go function can return multiple values
	return isValidName, isValidEmail, isValidTicketNumber
}
