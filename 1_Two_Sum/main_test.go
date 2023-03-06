package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	i := 1
	j := 1

	result := TwoSum(i, j)
	expectedResult := 2

	assert.Equal(t, expectedResult, result)
}
