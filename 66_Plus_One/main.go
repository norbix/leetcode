package main

import "fmt"

func PlusOne(digits []int) []int {
	n := len(digits)
	carry := 1 // Initialize the carry to 1 because we want to increment by one.

	for i := n - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}

	// If there is still a carry after the loop, we need to add a new most significant digit.
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits

}

func main() {
	input := []int{1, 2, 3}
	result := PlusOne(input)

	fmt.Printf("Output: %v", result)
}
