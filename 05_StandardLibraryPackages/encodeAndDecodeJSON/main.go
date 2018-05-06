package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID                         int
	FirstName, LastName, Title string
}

// employess2 uses struct tags to set the name of values in JSON
type employee2 struct {
	ID        int    `json:"Number"`
	FirstName string `json:"fName"`
	LastName  string `json:"lName"`
	Title     string `json:"Position"`
}

func main() {
	{ // Without struct tags
		e := employee{
			ID:        34,
			FirstName: "Auston",
			LastName:  "Matthews",
			Title:     "Center",
		}

		// Encode an employee into JSON
		data, err := json.Marshal(e)
		if err != nil {
			log.Print(err)
		}
		// Convert JSON bytes to a string for printing
		jsonString := string(data)
		fmt.Println("JSON string is ", jsonString)

		// Example []byte of JSON
		b := []byte(`{"ID":16,"FirstName":"Mitch","LastName":"Marner","Title":"Winger"}`)
		var e2 employee
		err = json.Unmarshal(b, &e2)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("ID: %d, Name: %s %s, JobTitle: %s\n\n", e2.ID, e2.FirstName, e2.LastName, e2.Title)
	}
	{ // With struct tags
		e := employee2{
			ID:        34,
			FirstName: "Auston",
			LastName:  "Matthews",
			Title:     "Center",
		}

		// Encode an employee into JSON
		data, err := json.Marshal(e)
		if err != nil {
			log.Print(err)
		}
		// Convert JSON bytes to a string for printing
		jsonString := string(data)
		fmt.Println("JSON string is ", jsonString)

		// Example []byte of JSON
		b := []byte(`{"number":16,"fName":"Mitch","lName":"Marner","position":"Winger"}`)
		var e2 employee2
		err = json.Unmarshal(b, &e2)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("ID: %d, Name: %s %s, JobTitle: %s", e2.ID, e2.FirstName, e2.LastName, e2.Title)
	}
}
