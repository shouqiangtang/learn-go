package main

import (
	"testing"
	"reflect"
)

func TestMerge(t *testing.T) {
	tests := []struct{
		arr1 []int
		arr2 []int
		expected []int
	}{
		{
			[]int{1,3,5},
			[]int{2,4,6},
			[]int{1,2,3,4,5,6},
		},
	}

	for _, test := range tests {
		actual := merge(test.arr1, test.arr2)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("mergeSort(%v, %v) = %v, but expected %v", test.arr1,
			test.arr2, actual, test.expected)
		}
	}
}
