package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	input := "III"
	expected := 3
	result := RomanToInt(input)

	assert.Equal(t, expected, result)
}
