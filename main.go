package main

import "fmt"

//Functions in Go
func concat(s1 string, s2 string) string {
	return s1 + s2
}

//You just can make a declaration of a variable in Go function after the parameters with they are the same type

func add ( x, y int) int {
	return x + y
}

func main() {
	// Declare a variable in hard way exemple in Go
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	// The simple way to declare a variable in Go
	congrats := "Congratulations! thats the simples way to declare a variable in Go"

	fmt.Println("That's a simple way to make a print in Go")
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
	fmt.Println(congrats)
	
	//Functions prints in Go
	fmt.Println(concat("Lane ", "Happy Birthday!"))
	fmt.Println(add(10, 20))

}