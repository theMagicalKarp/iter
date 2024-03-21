package predicates

import "strings"

// Contains returns a Predicate function that checks if a given string contains a specified substring.
func Contains(substr string) Predicate[string] {
	return func(s string) bool {
		return strings.Contains(s, substr)
	}
}

// HasPrefix returns a predicate function that checks if a string has the specified prefix.
// It takes a prefix string as input and returns a Predicate[string] function.
// The returned function checks if the input string starts with the specified prefix and returns true if it does,
// false otherwise.
func HasPrefix(prefix string) Predicate[string] {
	return func(s string) bool {
		return strings.HasPrefix(s, prefix)
	}
}

// HasSuffix returns a predicate function that checks if a string has the specified suffix.
// It takes a suffix string as input and returns a Predicate[string] function.
// The returned function takes a string as input and returns true if the string has the specified suffix,
// false otherwise.
func HasSuffix(suffix string) Predicate[string] {
	return func(s string) bool {
		return strings.HasSuffix(s, suffix)
	}
}
