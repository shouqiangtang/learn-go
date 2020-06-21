package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{4, 6, 3, 1, 2, 5}, expected: 3},
		{input: []int{4, 3, 1, 2, 6, 7, 8, 5}, expected: 3},
	}

	for _, test := range tests {
		actual := partition(&test.input, 0, len(test.input)-1)
		assert.Equal(t, test.expected, actual)
		t.Log(test.input)
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{4, 6, 3, 1, 2, 5}, expected: 3},
	}

	for _, test := range tests {
		QuickSort(&test.input, 0, len(test.input)-1)
		t.Log(test.input)
	}
}
