package main

import (
	"flag"
	"fmt"
)

func main() {
	fileName := flag.String("filename", "logfile", "File name for the log file")
	logLevel := flag.Int("loglevel", 0, "An integer value for the level of logging (0-4)")
	isEnable := flag.Bool("enable", false, "A boolean value for enabling log options")

	// Go also provides -h or --help which prints all flags as well as their usage definitions

	// Bind the flag to a variable
	var num int
	flag.IntVar(&num, "num", 25, "An integer value")

	// Parse definitions from the argument list
	flag.Parse()

	// Get the values from pointers
	fmt.Println("filename:", *fileName)
	fmt.Println("loglevel:", *logLevel)
	fmt.Println("enable:", *isEnable)

	// Get the value from a variable
	fmt.Println("num:", num)

	// Check for non-flag commandline arguments
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println("The non-flag command-line arguments are:")
		// Print the arguments
		for _, v := range args {
			fmt.Println(v)
		}
	}
}
