package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Function to check if a number is a palindrome
func IsPalindrome(n int) (bool, error) {
	// Throw an error if the number is negative
	if n < 0 {
		return false, errors.New("the number is negative, cannot check for palindrome")
	}

	// Convert the number to a string
	str := strconv.Itoa(n)

	// Check if the number is the same when read backwards
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false, nil
		}
	}

	// If the number reads the same forwards and backwards, it is a palindrome
	return true, nil
}

func main() {
	// Test examples
	numbers := []int{121, -121, 12321, 123, 0}

	for _, num := range numbers {
		isPal, err := IsPalindrome(num)
		if err != nil {
			fmt.Printf("Number %d: %v\n", num, err)
		} else if isPal {
			fmt.Printf("Number %d is a palindrome.\n", num)
		} else {
			fmt.Printf("Number %d is not a palindrome.\n", num)
		}
	}
}
