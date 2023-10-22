package tutorial

import (
	"log"
	"regexp"
	"strings"
)

func Regex() {
	log.Println("Start Regex tutorial")

	s1 := "peach"
	s2 := "peach punch" // for comparison
	s3 := "peach punch pinch"
	log.Printf("The string is '%s'", s1)

	// tests whether a pattern matches a string
	// we can use a string pattern directly
	match1, _ := regexp.MatchString("p([a-z]+)ch", s1)
	log.Println("The string matches the regex 'p([a-z]+)ch':", match1) // should be true

	// For other regexp tasks you’ll need to Compile an optimized Regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Match
	log.Println("The string matches the regex 'p([a-z]+)ch':", r.MatchString(s1)) // true

	// Substring
	// Only returns the first substring that matches, so 'punch' will not be returned
	substring := r.FindString(s2)
	if substring != "" {
		log.Println("Found substring:", substring)
	} else {
		log.Println("Substring not found")
	}

	// All substrings
	// use '-1' to find them all (and in the darkness, bind them ...)
	allSubstrings := r.FindAllString(s3, -1)
	if allSubstrings != nil {
		log.Println("Found substring:", allSubstrings)
	} else {
		log.Println("Substring not found")
	}
	twoFirstSubstrings := r.FindAllString(s3, 2)
	log.Println("Only the two first matches:", twoFirstSubstrings)

	// Replace
	log.Println("Replacing substring:", r.ReplaceAllString("peach drink", "Banana"))

	// Replace using a function
	// Better for string operations to perform on matches
	// Example : turn matches in capital letters
	log.Println("Matches to capital letters:", r.ReplaceAllStringFunc("peach drink", strings.ToUpper))

	// Substring position
	// returns a slice
	indexes := r.FindStringIndex(s2)
	log.Println("String found in [start, end] position:", indexes)

	// Substring submatch
	// The Submatch variants include information about both the whole-pattern matches and the submatches
	// within those matches.
	// For example this will return information for both p([a-z]+)ch and ([a-z]+).
	// You can also simply use FindSubmatch() for submatches only.
	submatch := r.FindStringSubmatch(s2)
	if submatch != nil {
		log.Println("Found match or submatch:", submatch)
	} else {
		log.Println("Found no match or submatch")
	}

	// Substring submatch position
	// Similarly this will return information about the indexes of matches and submatches.
	submatchIndexes := r.FindStringSubmatchIndex(s2)
	log.Println("Submatches found in [start, end] positions:", submatchIndexes)
}
