package running_sums

import (
	"errors"
	"fmt"
)

// returns a slice of sums for the parameter nums
func running_sums(nums []int) ([]int, error) {
	// If no nums were given, return an error with a message.
	if nums == nil {
		return []int, errors.New("empty nums")
	}

	// declare splice result
	var runningSums []int
		
	// loop through parameter splice and append sum to result
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		runningSums = append(runningSums, sum)
	}

	// Return a slice of sums
	return runningSums
}