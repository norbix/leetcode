package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	input := generatePositiveFixture()
	for _, v := range input {
		result := IsPalindrome(v)
		assert.True(t, result, true)
	}

	input = generateNegativeFixture()
	for _, v := range input {
		result := IsPalindrome(v)
		assert.False(t, result, false)
	}

}

func generatePositiveFixture() []int {
	return []int{121, 1331, 12321}
}

func generateNegativeFixture() []int {
	return []int{-121, -1331, -12321}
}
