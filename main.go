package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const ConferenceTickets = 20
	var remainingTickets = 20
	var bookings[] string 

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v and only %v tickets remaining\n", ConferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")

	for {

		var UserName string
	    var UserEmail string
	var UserTickets int

	// Get user details
	fmt.Println("Please enter your name: ")
	fmt.Scan(&UserName)
	fmt.Println("Please enter your email: ")
	fmt.Scan(&UserEmail)
	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(&UserTickets)
	
	remainingTickets = remainingTickets - UserTickets
	bookings = append(bookings, UserName)


	fmt.Printf("Thank you %v for booking %v tickets , you will receive a notification on %v , please check your inbox!!!\n", UserName, UserTickets, UserEmail)
	fmt.Printf("We have %v tickets remaining\n", remainingTickets)


	fmt.Printf("Bookings: %v\n", bookings)

	}

	


	

}


