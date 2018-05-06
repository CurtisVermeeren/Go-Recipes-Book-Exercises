package main

import (
	"fmt"
)

// user defined struct
type customer struct {
	FirstName string
	LastName  string
	Age       int
}

// user defined types as type for a field
type order struct {
	item   string
	price  float32
	Person customer
}

// embedding customer type into human
type human struct {
	customer
	isHuman bool
}

// adding method to struct
func (c *customer) toString() string {
	return fmt.Sprintf("%s %s is %d years old", c.FirstName, c.LastName, c.Age)
}

// use pointer to customer to change a field in a method
func (c *customer) getOld() {
	c.Age++
}

func main() {
	{ // Creating an instance of a struct
		var c customer
		c.FirstName = "Curtis"
		c.LastName = "Vermeeren"
		c.Age = 24

		c2 := &customer{
			FirstName: "Auston",
			LastName:  "Matthews",
			Age:       20,
		}

		fmt.Println(c)
		fmt.Println(c2)

		fmt.Println(c.toString())
		c.getOld()
		fmt.Println(c.toString())

		o := &order{
			item:   "Hyken Chair",
			price:  184.99,
			Person: c,
		}
		fmt.Println(o)
	}
}
