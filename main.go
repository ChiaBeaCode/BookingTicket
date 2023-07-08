package main

import (
	"BookingTicket/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

// var bookings = []string{}
var bookings = make([]userDataa, 0)

type userDataa struct {
	firstName        string
	lastName         string
	email            string
	ticketsPurchased uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()
	firstName, lastName, email, userTickets := userEntries()

	// bookings = append(bookings, firstName+" "+lastName)

	isValidName, isValidEmail, isValidUserTickers := helper.ValidateUserEntries(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidUserTickers {
		bookedTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// fmt.Printf("Remaining tickets: %v\n", remainingTickets)
		if remainingTickets == 0 {
			fmt.Println("Conference booked")
			fmt.Printf("Guests: %v\n", printGuestFirstNames())
		}

	} else {
		if !isValidName {
			fmt.Println("First and last name must exceed 2 characters, try again")
		}
		if !isValidEmail {
			fmt.Println("Please enter a valid email")
		}
		if !isValidUserTickers {
			fmt.Printf("You entered: %v. We have %v remaining tickets, try again\n", userTickets, remainingTickets)
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printGuestFirstNames() []string {
	firstNames := []string{}
	for _, name := range bookings {
		firstNames = append(firstNames, name.firstName)
	}
	return firstNames
}

func userEntries() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter email: ")
	fmt.Scan(&email)

	fmt.Println("Enter amount of tickets you wish to purchase: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets

}

func bookedTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	//automatically stores last name in lastName when both entered on firstName input
	// bookings = append(bookings, firstName+" "+lastName)
	var userData = userDataa{
		firstName:        firstName,
		lastName:         lastName,
		email:            email,
		ticketsPurchased: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for purchasing %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("Your purchase confirmation with be sent to your email at: %v\n", email)
	// sendTicket(userTickets, firstName, lastName, email)
	// fmt.Printf("Type: %T\nValue: %v\n", userData, bookings)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	//helps put 2gether a string in formatted output that allow for storage in strvar, unlike Printf
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########################")
	fmt.Printf("Sending ticket:\n%v\nto email address %v\n", ticket, email)
	fmt.Println("##########################")
	wg.Done()

}
