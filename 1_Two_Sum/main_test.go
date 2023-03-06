package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum2(t *testing.T) {
	t.Skip()

	i := 3
	j := 3

	result := TwoSum(i, j)
	expectedResult := 6

	assert.Equal(t, expectedResult, result)
}

func TestTwoSum(t *testing.T) {
	i := 1
	j := 1

	result := TwoSum(i, j)
	expectedResult := 2

	assert.Equal(t, expectedResult, result)
}
