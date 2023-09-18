package main

import "fmt"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	shortest := strs[0]
	for _, str := range strs {
		if len(str) < len(shortest) {
			shortest = str
		}
	}

	// Iterate through the characters of the shortest string.
	for i, char := range shortest {
		for _, str := range strs {
			if str[i] != byte(char) {
				// If a character doesn't match, return the prefix up to this point.
				return shortest[:i]
			}
		}
	}

	return shortest
}

func main() {
	input := []string{"flower", "flow", "flight"}
	lCP := LongestCommonPrefix(input)

	fmt.Printf("The common prefix is %v", lCP)
}
