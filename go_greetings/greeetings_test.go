package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with name,
// checking if a valid greeting is returned.
func TestHelloName(t *testing.T) {
	name := "Luke"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := greetings.Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Luke") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with empty string,
// checking for an error
func TestHelloEmpty(t *testing.T) {
	msg, err := greetings.Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
