package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	contents, err := readFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(contents))
	fmt.Println("\nStarting Panic!")
	panicRecover()
	fmt.Println("Program has regained control!")
}

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Defer closing the file till the end of the function

	return ioutil.ReadAll(file)
}

func panicRecover() {
	defer fmt.Println("1st Defer Call")
	defer func() {
		fmt.Println("2nd Defer Call")
		if e := recover(); e != nil {
			fmt.Println("Recover with: ", e)
		}
	}()
	panic("Panicing to show panic!")
	fmt.Println("This is never called")
}
