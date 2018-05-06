package main

import "fmt"

func main() {
	{ // Declare and initialize a slice
		var x []int
		x = make([]int, 3, 6)
		x[0] = 5
		x[1] = 10
		x[2] = 15
		fmt.Println(x)
	}

	{ // Slice literals
		x := []int{4, 6, 8}
		fmt.Println(x)

		z := []int{0: 3, 2: 9, 5: 18}
		fmt.Println(z)

		y := []int{} // empty slice

		var q []int // nil slice

		fmt.Println(y == nil)
		fmt.Println(q == nil)
	}

	{ // Larger slice by copying
		x := []int{10, 20, 30}
		fmt.Printf("[Slice:x] Length is %d Capacity is %d\n", len(x), cap(x))
		// Create a bigger slice
		y := make([]int, 5, 10)
		copy(y, x)
		fmt.Printf("[Slice:y] Length is %d Capacity is %d\n", len(y), cap(y))
		fmt.Println("Slice y after copying:", y)
		y[3] = 40
		y[4] = 50
		fmt.Printf("Slice y after adding elements: %d\n", y)
	}

	{ // Appending slices
		x := make([]int, 2, 5)
		x[0] = 10
		x[1] = 20
		fmt.Println("Slice x:", x)
		fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x))
		// Create a bigger slice
		x = append(x, 30, 40, 50)
		fmt.Println("Slice x after appending data:", x)
		fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x))
		x = append(x, 60, 70, 80)
		fmt.Println("Slice x after appending data for the second time:", x)
		fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x))
	}

	{ // Iterating over a slice
		x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for k, v := range x {
			fmt.Printf("x[%d]: %d\n", k, v)
		}

	}
}
