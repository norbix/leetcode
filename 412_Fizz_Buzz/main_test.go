package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	input := 3
	expected := []string{"1", "2", "Fizz"}
	result := FizzBuzz(input)
	assert.Equal(t, expected, result)

	input = 5
	expected = []string{"1", "2", "Fizz", "4", "Buzz"}
	result = FizzBuzz(input)
	assert.Equal(t, expected, result)

	input = 15
	expected = []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	result = FizzBuzz(input)
	assert.Equal(t, expected, result)
}
