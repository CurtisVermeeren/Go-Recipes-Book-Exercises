package main

import "fmt"

// Variadic Function parameters
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Passing functions as values
func splitValues(f func(sum int) (int, int)) {
	x, y := f(35)
	fmt.Println(x, y)
	x, y = f(50)
	fmt.Println(x, y)
}

func main() {
	{ // Variadic Functions
		fmt.Println(sum(1, 2, 3, 4, 5, 6))
		fmt.Println(sum(100, 2, 4, 6))
		fmt.Println(sum(1))
	}

	// Anonymous function
	a, b := 5, 7
	fn := func(sum int) (int, int) {
		x := sum * a / b
		y := sum - x
		return x, y
	}

	{ // Passing functions as values
		splitValues(fn)
	}
}
