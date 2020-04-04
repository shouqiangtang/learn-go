package nonrepeating

import (
	"testing"
)

func TestMaxLengthOfNoRepeatingSubStr(t *testing.T) {
	tests := []struct {
		input string
		len   int
	}{
		{"aaaaa", 1},
		{"abcab", 3},
		{"abcabcde", 5},
	}

	for _, test := range tests {
		actual := MaxLengthOfNoRepeatingSubStr(test.input)
		if actual != test.len {
			t.Errorf("MaxLengthOfNoRepeatingSubStr(%s) = %d, but has %d",
				test.input, test.len, actual)
		}
	}
}

func TestCheckRepeat(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", false},
		{"a", false},
		{"aaaa", true},
		{"abcde", false},
	}

	for _, test := range tests {
		actual := checkRepeat(test.input)
		// fmt.Printf("%#v\n", actual != test.expected)
		if actual != test.expected {
			t.Errorf("input: %s expected: %v, but has %v",
				test.input, test.expected, actual)
		}
	}
}

func TestMaxNoRepeatingSubStr(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"aaaaa", "a"},
		{"abcabcde", "abcde"},
	}

	for _, test := range tests {
		actual := MaxNoRepeatingSubStr(test.input)
		if actual != test.expected {
			t.Errorf("input %s, expected %s, actual %s",
				test.input, test.expected, actual)
		}
	}
}
