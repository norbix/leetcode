package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	input := []string{"flower", "flow", "flight"}
	expected := "fl"
	result := LongestCommonPrefix(input)

	assert.Equal(t, expected, result)
}
