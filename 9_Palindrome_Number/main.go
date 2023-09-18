package main

import "fmt"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false // Negative numbers are not palindromes.
	}

	// Reverse the integer.
	reversed, original := 0, x
	for original > 0 {
		digit := original % 10
		reversed = reversed*10 + digit
		original /= 10
	}

	return reversed == x
}

func main() {
	input := 121

	if IsPalindrome(input) {
		fmt.Println("The integer is a palindrome.")
	} else {
		fmt.Println("The integer is not a palindrome.")
	}
}
