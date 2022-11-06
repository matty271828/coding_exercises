package main

import "fmt"

func main() {
	// Question 1
	//fmt.Println(check_unique("lola"))

	// Question 2
	fmt.Println(check_permutation("dog", "god"))
}

// Implement an algo to determine if a string has all unique characters
// What if you cannot you additional data structures
func check_unique(check_string string) bool {
	// Create map of bool values for whether a char found
	isPresent := make(map[byte]bool)
	for i := range check_string {
		isPresent[check_string[i]] = false
	}

	// Loop through each char in string
	for i := range check_string {
		// Call key from map, if already present exit loop and return False
		if isPresent[check_string[i]] == true {
			return false
		} else {
			// Else set value to true
			isPresent[check_string[i]] = true
		}
	}
	// All characters checked - return true
	return true
	// TIME COMPLEXITY -> O(n)
	// SPACE COMPLEXITY -> O(1) (only need to store bool values for maximum 128 characters)
}

// Given two strings, write a method to decide if one is a permutation of the other
func check_permutation(string1, string2 string) bool {
	// Permutation exists if strings consist of all the same characters

	// Create array to store int value of each value in string
	var count [128]int

	// Count number of occurences of each letter in string 1
	// Iterate through string and +1 to value for each byte
	for i := range string1 {
		count[string1[i]]++
	}

	// Count number of occurences of each letter in string 2
	for i := range string2 {
		// Iterate through string and -1 to value for each byte
		count[string1[i]]--
	}

	// Loop through count, return false if any value non-zero
	for i := range count {
		if count[i] != 0 {
			return false
		}
	}
	return true

	// TIME COMPLEXITY -> O(n)
	// SPACE COMPLEXITY -> O(1)
}
