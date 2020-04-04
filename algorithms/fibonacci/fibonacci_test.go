package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibo(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{10, 55},
	}

	for _, test := range tests {
		actual := Fib(test.input)
		if actual != test.expected {
			t.Errorf(
				"Fibo(%d) = %d, but has %d",
				test.input, test.expected, actual)
		}
	}
}

func TestFiboBuf(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{10, 55},
	}

	for _, test := range tests {
		actual := FibBuf(test.input)
		if actual != test.expected {
			t.Errorf(
				"FiboBuf(%d) = %d, but has %d",
				test.input, test.expected, actual)
		}
	}

	t.Logf("k = %d", k)
}

func TestLocateFibo(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1, 2},
		{2, 4},
		{3, 5},
		{5, 6},
		{34, 10},
		{12, -1},
	}

	for _, test := range tests {
		actual := LocateFib(test.input)
		if actual != test.expected {
			t.Errorf("LocateFibo(%d) = %d, but has %d",
				test.input, test.expected, actual)
		}
	}
}

func TestSumFibo(t *testing.T) {
	// 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
	tests := []struct {
		input    int
		expected int
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 4},
		{10, 88},
	}

	for _, test := range tests {
		actual := SumFib(test.input)
		if actual != test.expected {
			t.Errorf("SumFibo(%d) = %d, but has %d",
				test.input, test.expected, actual)
		}
	}
}

func TestFiboList(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{1, []int{0}},
		{2, []int{0, 1}},
		{3, []int{0, 1, 1}},
		{4, []int{0, 1, 1, 2}},
		{9, []int{0, 1, 1, 2, 3, 5, 8, 13, 21}},
	}

	for _, test := range tests {
		actual := FibList(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("FiboList(%d) = %v, but has %v",
				test.input, test.expected, actual)
		}
	}
}

func TestFibLoop(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{10, 55},
	}

	for _, test := range tests {
		actual := FibLoop(test.input)
		if actual != test.expected {
			t.Errorf("FibLoop(%d) = %d, but has %d",
				test.input, test.expected, actual)
		}
	}
}
