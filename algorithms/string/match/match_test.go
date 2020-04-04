package match

import "testing"

func TestIndex(t *testing.T) {
	tests := []struct {
		str, subStr string
		expected    int
	}{
		{str: "abc", subStr: "a", expected: 1},
		{str: "abcadcabcda", subStr: "abcd", expected: 7},
		{str: "唐守强唐可凡唐静涵张春焕", subStr: "可凡", expected: 5},
	}

	for _, test := range tests {
		actual := Index(test.str, test.subStr, 1)
		if actual != test.expected {
			t.Errorf("Index(%s, %s, 0) = %d, but expected %d",
				test.str, test.subStr, actual, test.expected)
		}
	}
}
