package main

import "fmt"

type teamMember interface {
	PrintName()
	PrintDetails()
}

type pet struct {
	name string
	age  int
}

type person struct {
	pet
	fName, lName string
	age          int
}

func (p pet) PrintName() {
	fmt.Println("Hello I'm", p.name)
}

// Person has methods to satisfy teamMember interface
func (p person) PrintName() {
	fmt.Println("Hello I'm", p.fName)
}

func (p person) PrintDetails() {
	fmt.Println(p.fName, p.lName, "is", p.age, "years old")
	fmt.Println(p.fName, "has a pet named", p.pet.name, "it is", p.pet.age, "years old")
}

// Use teamMember as a type and call it's methods
func printTeamMember(t teamMember) {
	t.PrintName()
	t.PrintDetails()
}

func main() {
	pt := pet{
		"Fluffy",
		12,
	}
	p := &person{
		pt,
		"Curtis",
		"Vermeeren",
		24,
	}
	// Using person as a teamMember
	printTeamMember(p)

	// Calling PrintName of person
	p.PrintName()
	// Calling PrintName of pet
	p.pet.PrintName()
}
