package main

import "fmt"

func main() {
	{ // Using var keyword
		var firstName string
		var lastName string
		var age int
		fmt.Println("First name: ", firstName)
		fmt.Println("Last name: ", lastName)
		fmt.Println("Age: ", age)
	}

	{ // Multiple values in a single statement
		var firstName, lastName string = "Curtis", "Vermeeren"
		age := 24
		fmt.Println("First name: ", firstName)
		fmt.Println("Last name: ", lastName)
		fmt.Println("Age: ", age)
	}

	{ // Short variable declaration using :=
		firstName, lastName := "Curtis", "Vermeeren"
		age := 24
		fmt.Println("First name: ", firstName)
		fmt.Println("Last name: ", lastName)
		fmt.Println("Age: ", age)
	}

}
