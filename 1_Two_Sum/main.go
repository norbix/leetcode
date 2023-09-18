package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			// We found a complement in the map, return the indices.
			return []int{index, i}
		}

		// Store the current number and its index in the map.
		numMap[num] = i
	}

	// If no solution is found, return an empty slice.
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	sum := TwoSum(nums, target)
	fmt.Printf("Indices of the two numbers of sum target are %v", sum)
}
