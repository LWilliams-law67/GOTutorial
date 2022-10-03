package running_sums

import (
	"regexp"
	"testing"
)

// TestRunningSums calls running_sums.running_sums with nums,
// checking if a valid slice of sums is returned.
func TestRunningSums(t *testing.T) {
	nums := [6]int{2,4,6,1,3,5}
	got,err := running_sums(nums)
	want := [6]int[2,6,12,13,16,21]

	if want != got || err != nil {
		t.Fatalf(`running_sums("[6]int{2,4,6,1,3,5}") = %q, %v, want match for %#q, nil`, got, err, want)
	}
}