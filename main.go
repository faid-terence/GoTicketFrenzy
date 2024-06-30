package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const ConferenceTickets = 20
	var remainingTickets = 20

	fmt.Println("Welcome to the", conferenceName, "booking application")
	fmt.Println("We have a total of ", ConferenceTickets, "and only", remainingTickets, "tickets remaining")
	fmt.Println("Get your tickets now!")


}


