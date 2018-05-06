package main

import (
	"errors"
	"fmt"
)

func main() {
	err := returnsErrorIfLargerInt(200)
	if err != nil {
		fmt.Println("there was an error:", err)
	}

	err = returnsErrorIfLargerInt(5)
	if err != nil {
		fmt.Println("there was an error:", err)
	}
}

func returnsErrorIfLargerInt(a int) error {
	if a > 100 {
		return errors.New("int is too large")
	}
	return nil
}
