package main

import (
	"fmt"
	"log"
	"leetcode/running_sums"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("running sum: ")
	log.SetFlags(0)

	// A slice of numbers to get running sums for
	nums := []int{2, 6, 3, 1}

	// Request running sums result
	sums, err := running_sums.running_sums(nums)

	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error, prints splice of sums
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
