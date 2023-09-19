package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 4}
	result := PlusOne(input)

	assert.Equal(t, expected, result)
}
