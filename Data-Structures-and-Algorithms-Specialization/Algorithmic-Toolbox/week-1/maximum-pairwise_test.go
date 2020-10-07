package course

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMaxPairwiseProductBrutal(t *testing.T) {
	arr := []int{2, 5, 8, 10, 11, 4, 14, 1}
	result := maxPairwiseProductBrutal(arr)

	assert.Equal(t, result, 154)
}

func TestMaxPairwiseProduct(t *testing.T) {
	arr := []int{2, 5, 8, 10, 11, 4, 14, 1}
	result := maxPairwiseProduct(arr)

	assert.Equal(t, result, 154)
}