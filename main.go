package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const ConferenceTickets = 20
	var remainingTickets = 20

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v and only %v tickets remaining\n", ConferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")


	var UserName string
	var UserTickets int

	// Get user details
	fmt.Println("Please enter your name: ")
	fmt.Scan(&UserName)
	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(&UserTickets)

	fmt.Printf("Thank you %v for booking %v tickets\n", UserName, UserTickets)

}


