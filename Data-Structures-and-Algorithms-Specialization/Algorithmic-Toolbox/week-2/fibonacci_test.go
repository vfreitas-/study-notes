package course

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFibonacciTable(t *testing.T) {
	result := fibonacciTable(10)

	assert.Equal(t, result, 34)
}