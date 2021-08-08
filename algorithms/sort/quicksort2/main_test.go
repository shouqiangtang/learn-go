package main

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct{
		input []int
		expected int
	}{
		{
			input: []int{5,1,2,6,8},
			expected: 2,
		},
	}

	for _, test := range tests {
		actual := partition(test.input, 0, len(test.input)-1)
		if actual != test.expected {
			t.Errorf("partition(%v, %d, %d) = %d, but actual %d", test.input, 0, len(test.input)-1, test.expected, actual)
		}
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct{
		input []int
		expected []int
	}{
		{
			input: []int{12,19,1,6,4,3,2,1,2},
			expected: []int{1,1,2,2,3,4,6,12,19},
		},
	}

	for _, test := range tests {
		quickSort(test.input, 0, len(test.input)-1)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("quickSort(%v, %d, %d), expected %v, but not", test.input, 0, len(test.input)-1, test.expected)
		}
	}
}
