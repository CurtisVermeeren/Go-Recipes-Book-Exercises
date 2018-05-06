package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// WaitGroup is used to wait for the program to finish goroutines
var wg sync.WaitGroup

func main() {
	// GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously
	runtime.GOMAXPROCS(1)

	// Create a buffered channel
	text := make(chan string, 20)

	// Add 2 goroutines to wait for
	wg.Add(2)
	fmt.Println("Starting Goroutines")
	go addTable(text)
	go multiTable(text)
	fmt.Println("Waiting to finish")

	// Wait here for all goroutines to finish
	go func(s chan string) {
		wg.Wait()
		// Close the channel once all have sent data
		close(s)
	}(text)

	// Print the values send to the channel
	printText(text)

	fmt.Println("Program Ending...")
}

func addTable(s chan string) {
	// Tell waitgroup a routine has ended
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		lineOutput := fmt.Sprintf("Addition Table for: %d", i)
		finalOutput := ""
		for j := 1; j <= 10; j++ {
			finalOutput = fmt.Sprintf("%s%d+%d=%d ", finalOutput, i, j, i+j)
		}
		s <- fmt.Sprintf("%s\n%s\n\n", lineOutput, finalOutput)
	}
}

func multiTable(s chan string) {
	// Tell waitgroup a routine has ended
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		lineOutput := fmt.Sprintf("Multiplication Table for: %d", i)
		finalOutput := ""
		for j := 1; j <= 10; j++ {
			finalOutput = fmt.Sprintf("%s%d+%d=%d ", finalOutput, i, j, i*j)
		}
		s <- fmt.Sprintf("%s\n%s\n\n", lineOutput, finalOutput)
	}
}

func printText(s chan string) {
	// Print all values in the buffered channel s
	for val := range s {
		fmt.Print(val)
	}
}
