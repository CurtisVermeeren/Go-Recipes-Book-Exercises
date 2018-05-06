package main

import (
	"fmt"
)

func main() {
	{ // Declare and init an array
		var x [5]int
		fmt.Println(x)

		y := [5]int{5, 10, 15, 20, 25}
		fmt.Println(y)

		langs := [4]string{0: "C++", 2: "Rust", 3: "Golang"}
		fmt.Println(langs[3])

		z := [...]int{10, 20, 30, 40, 50, 60} // Can use [...] when you declare and init an array at the same time
		fmt.Println("Value of z:", z)
		fmt.Println("Length of z:", len(z))
	}

	{ // Iterating an array
		langs := [4]string{"Go", "Rust", "Scala", "Julia"}
		for i := 0; i < len(langs); i++ {
			fmt.Print(langs[i])
		}
		fmt.Println()

		for k, v := range langs { // k is index, v is value
			fmt.Print(k, ":", v, " ")
		}
	}
}
