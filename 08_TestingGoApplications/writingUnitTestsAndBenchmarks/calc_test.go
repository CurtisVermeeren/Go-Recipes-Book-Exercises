package calc

import (
	"fmt"
	"testing"
	"time"
)

// Test case for the Sum function
func TestSum(t *testing.T) {
	input, expected := []int{7, 8, 10}, 25
	result := Sum(input...)
	if result != expected {
		t.Errorf("Result: %d, Expected: %d", result, expected)
	}
}

// Test case for the Sum function
func TestAverage(t *testing.T) {
	input, expected := []int{7, 8, 10}, 8.33
	result := Average(input...)
	if result != expected {
		t.Errorf("Result: %f, Expected: %f", result, expected)
	}
}

/*
* Tests can be run from console using "go test"
* option "-v" gives descriptive information when running tests
* option  "-cover" gets a percentage of tests casts written vs the code
* option "-short" signals to skip long running tests
* 	short can be checks using testing.Short(), then using t.Skip("Message")
 */

// TestLongRun is a time-consuming test
func TestLongRun(t *testing.T) {
	// Checks whether the short flag is provided
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
	// Long running implementation goes here
	time.Sleep(5 * time.Second)
}

/*
* Benchmarks can be run using "go test -v -bench=."
* or "go test -v -bench ."
* The tests will be run b.N times till a reliable ping is reached
 */

// Benchmark for function Sum
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(7, 8, 10)
	}
}

// Benchmark for function Average
func BenchmarkAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Average(7, 8, 10)
	}
}

/*
* Go tests can be run in parallel to save time
* call the method t.Parallel() for any function to run in parallel when testing
* t.Parallel() should be the first statement called
 */

// Test case for the function Sum to be executed in parallel
func TestSumInParallel(t *testing.T) {
	t.Parallel()
	// Delaying 1 second for the sake of demonstration
	time.Sleep(1 * time.Second)
	input, expected := []int{7, 8, 10}, 25
	result := Sum(input...)
	if result != expected {
		t.Errorf("Result: %d, Expected: %d", result, expected)
	}
}

// Test case for the function Sum to be executed in parallel
func TestAverageInParallel(t *testing.T) {
	t.Parallel()
	// Delaying 1 second for the sake of demonstration
	time.Sleep(2 * time.Second)
	input, expected := []int{7, 8, 10}, 8.33
	result := Average(input...)
	if result != expected {
		t.Errorf("Result: %f, Expected: %f", result, expected)
	}
}

/*
* exmaple functions can be written by using the prefix "Example" in the function name
* func Example() // Example test for package
* func ExampleF() // Example test for function F
* func ExampleT() // Example test for type T
* func ExampleT_M() // Example test for M on type T
*
* You typically include a concluding line comment that begins with Output:
* Compares the given output with the output function when the tests are executed
 */

// Example code for function Sum
func ExampleSum() {
	fmt.Println(Sum(7, 8, 10))
	// Output: 25
}

// Example code for function Average
func ExampleAverage() {
	fmt.Println(Average(7, 8, 10))
	// Output: 8.33
}
