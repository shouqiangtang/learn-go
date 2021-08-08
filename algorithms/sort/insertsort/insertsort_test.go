package insertsort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 3, 2, 6, 7, 1}, expected: []int{1, 2, 3, 5, 6, 7}},
		{input: []int{12, 11, 10, 9, 8}, expected: []int{8, 9, 10, 11, 12}},
		{input: []int{1}, expected: []int{1}},
		{input: []int{}, expected: []int{}},
	}

	for _, test := range tests {
		actual := InsertSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("InsertSort(%v) = %v, but has %v",
				test.input, test.expected, actual)
		}
	}
}
