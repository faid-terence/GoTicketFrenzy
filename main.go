package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const ConferenceTickets = 20
	var remainingTickets = 20
	var bookings[] string 

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v and only %v tickets remaining\n", ConferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")

	for {

		var firstName string
		var lastName string
	    var UserEmail string
	var UserTickets int

	// Get user details
	fmt.Println("Please enter your name: ")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email: ")
	fmt.Scan(&UserEmail)
	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(&UserTickets)

	if UserTickets <= remainingTickets {
		fmt.Printf("Sorry, we only have %v tickets remaining\n", remainingTickets)
		remainingTickets = remainingTickets - UserTickets
		bookings = append(bookings, firstName + " " + lastName)
	
	
		fmt.Printf("Thank you %v for booking %v tickets , you will receive a notification on %v , please check your inbox!!!\n", lastName, UserTickets, UserEmail)
		fmt.Printf("We have %v tickets remaining\n", remainingTickets)
	
		firstNames:= []string {}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
	
		fmt.Printf("First names of people who have booked: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Sorry, we are sold out!")
			break
		}
	}else {
		fmt.Printf("Sorry, we only have %v tickets remaining\n", remainingTickets)
	}
	



	}

	


	

}


