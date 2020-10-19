package main

import (
	"testing"
	"reflect"
)

func TestMergeElems(t *testing.T) {
	tests := []struct {
		input []int
		expected []int
	}{
		{
			input: []int{1,2,3,4,5,6,7,8},
			expected: []int{3,3,6,9,15},
		},
		{
			input: []int{1,1,2,2,1,1,3,4,5,6,19},
			expected: []int{3,3,3,6,6,24},
		},
		{
			input: []int{123, 34, 78, 88, 29, 10, 34, 76},
			expected: []int{123, 78, 39, 144, 88},
		},
	}

	for _, test := range tests {
		actual := mergeElems(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("expected: %v, but has %v\n", test.expected, actual)
		}
	}
}
