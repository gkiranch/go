package main

import "fmt"

func main() {
	var conferenceName = "Go Conference" // we use var but can use constant
	const conferenceTickets = 50         // we are using constant value that doesn't change
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your booking tickets here")

	var userName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your first name")
	fmt.Scan(&userName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email at %v\n", userName, lastName, userTickets, email)
	fmt.Printf("We have total of %v tickets remaining for %v\n", remainingTickets, conferenceName)
}
