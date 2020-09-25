package search

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearchNumber(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarySearchNumber(arr, 2)

	assert.Equal(t, result, 1)
}

func TestBinarySearchString(t *testing.T) {
	arr := []string{"C++", "Elixir", "Golang", "Javascript", "NodeJS", "PHP", "Rust"}
	result := binarySearchString(arr, "Golang")

	assert.Equal(t, result, 2)
}

