package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type fibvalue struct {
	input, value int
}

var wg sync.WaitGroup

func main() {
	// Seed Rand with the time
	rand.Seed(time.Now().UnixNano())

	// Add 3 to waitgroup
	wg.Add(3)
	// Declare channels
	randoms := make(chan int)
	fibs := make(chan fibvalue)

	// Launch goroutines
	go randomCounter(randoms)
	go generateFinonacci(fibs, randoms)
	go printFibonacci(fibs)

	// Wait for completion
	wg.Wait()
}

// randomCounter generates random values
func randomCounter(out chan int) {
	defer wg.Done()
	var random int
	// Send 10 random values to out channel
	for x := 0; x < 10; x++ {
		random = rand.Intn(50)
		out <- random
	}
	close(out)
}

// generateFinonacci produces the fib values of inputs from in channel
func generateFinonacci(out chan fibvalue, in chan int) {
	defer wg.Done()
	var input float64
	for v := range in {
		input = float64(v)
		Phi := (1 + math.Sqrt(5)) / 2
		phi := (1 - math.Sqrt(5)) / 2
		result := (math.Pow(Phi, input) - math.Pow(phi, input)) / math.Sqrt(5)
		out <- fibvalue{
			input: v,
			value: int(result),
		}
	}
	close(out)
}

// printFibonacci prints the value received on the in channel
func printFibonacci(in chan fibvalue) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("Fibonacci value of %d is %d\n", v.input, v.value)
	}
}
