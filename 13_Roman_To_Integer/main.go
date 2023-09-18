package main

import "fmt"

func RomanToInt(s string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prevValue := 0 // Keep track of the previous symbol's value.

	for i := len(s) - 1; i >= 0; i-- {
		current := romanValues[s[i]]
		if current < prevValue {
			result -= current // Subtract if the current value is smaller than the previous.
		} else {
			result += current // Otherwise, add the current value.
		}
		prevValue = current
	}

	return result
}

func main() {
	input := "III"
	result := RomanToInt(input)

	fmt.Printf("Roman number %v, converted to an int %v\n", input, result)
}
