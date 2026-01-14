package main

import "fmt"

func main() {
	//-------------------------------------------------------------------------------------------
	//Declaring a variable
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username,)
	//-------------------------------------------------------------------------------------------
	//Short variable declaration
	congrats := "happy birthday!"

	fmt.Println(congrats)
	//-------------------------------------------------------------------------------------------
	//Same line declaration
	averageOpenRate, displayMessage := .23, "is the average open rate of your message"

	fmt.Println(averageOpenRate, displayMessage)
	//-------------------------------------------------------------------------------------------
	//Converting Between Types
	accountAge := 2.6

	accountAgeInt := int(accountAge)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
	//-------------------------------------------------------------------------------------------
	//Consts
	const premiumPlaneName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("plan:", premiumPlaneName)
	fmt.Println("plan:", basicPlanName)
	//-------------------------------------------------------------------------------------------
	//Computed Constants
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("number of seconds in an hour:", secondsInHour)
	//-------------------------------------------------------------------------------------------
	//Formatting strings
	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	fmt.Println(msg)	
	//-------------------------------------------------------------------------------------------
	//Conditionals
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}