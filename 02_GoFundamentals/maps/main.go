package main

import (
	"fmt"
	"sort"
)

func main() {
	{ // Declare and intitialize maps
		var chapts map[int]string
		chapts = make(map[int]string)

		chapts[1] = "Introduction"
		chapts[2] = "Begin Go"
		chapts[3] = "Structs and interfaces"
		chapts[4] = "Databases"

		fmt.Println(chapts)

		langs := map[string]string{
			"EN": "English",
			"ES": "Spanish",
			"FR": "French",
		}

		fmt.Println(langs)

		// Working with maps

		if lang, ok := langs["EN"]; ok { // ok == true if key exists
			fmt.Println(lang)
		} else {
			fmt.Println("Doesn't exist")
		}

		delete(langs, "EN") // Delete EN from langs

		if lang, ok := langs["EN"]; ok {
			fmt.Println(lang)
		} else {
			fmt.Println("Doesn't exist")
		}

		// slice for order of keys in chapts map
		var keys []int
		for key := range chapts {
			keys = append(keys, key)
		}
		sort.Ints(keys) // Sort keys in increasing order

		for _, key := range keys {
			fmt.Printf("Key: %d Value: %s \n", key, chapts[key])
		}
	}
}
