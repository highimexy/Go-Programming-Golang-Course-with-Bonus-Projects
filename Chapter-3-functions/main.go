package main

import "fmt"

//----------------------------------------------------------------
//Functions
func concat(s1 string, s2 string) string {
	return s1 + s2
}

//----------------------------------------------------------------
//Multiple Parameters
//func concat(s1, s2 string) string {
//	return s1 + s2
// }

func main(){
	fmt.Println(concat("Lane,", " happy birthday!"))
	fmt.Println(concat("Elon,", " hope that Tesla thing works out"))
	fmt.Println(concat("Go", " is fantastic"))
	//Passing Variables by value - not doing shit
	x := 5
	increment(x)

	fmt.Println(x)

	//Passing Variables by value - doing shit
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	
	fmt.Println("you`ve sent", sendsSoFar, "messages")
	//Ignoring return values
	// _
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
}

func increment (x int) {
	x++
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = (sendsSoFar + sendsToAdd)
	return sendsSoFar
}

func getNames() (string, string){
	return "John", "Doe"
}