package main

import (
	"fmt"
)

func main() {
	var x = "BOOKING"
	fmt.Println("WELCOME TO TICKETS", x)
	fmt.Println("GET YOUR TICKETS HERE")
	var remainingTick uint = 50
	fmt.Println("Remaining tickets =", remainingTick)
	var bookings []string
	// for {
	var Firstname string
	var Lastname string
	var userTick uint
	var email string
	var password string
	fmt.Println("Enter Firstname:")
	fmt.Scan(&Firstname)
	fmt.Println("Enter Lastname:")
	fmt.Scan(&Lastname)
	fmt.Println("Enter email id:")
	fmt.Scan(&email)
	fmt.Println("Enter password:")
	fmt.Scan(&password)
	fmt.Println("Enter no of tickets to book:")
	fmt.Scan(&userTick)
	fmt.Printf("Thank You %v %v for Booking %v Tickets .A confimation mail has been send to your mail:%v\n", Firstname, Lastname, userTick, email)
	remainingTick = remainingTick - userTick
	fmt.Printf("Remaining Tickets =%v\n", remainingTick)
	bookings = append(bookings, Firstname+" "+Lastname)
	fmt.Printf("All bookings in application %v\n", bookings)
}
