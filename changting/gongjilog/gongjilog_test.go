package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct{
		lines []LogLine
		ts int64
		flag int
		expected int
	}{
		{
			lines: []LogLine{
				LogLine{
					Ts: 1,
					Typ: "sqli",
					Cnt: 2,
				},
				LogLine{
					Ts: 3,
					Typ: "sqli",
					Cnt: 2,
				},
				LogLine{
					Ts: 5,
					Typ: "sqli",
					Cnt: 2,
				},
				LogLine{
					Ts: 7,
					Typ: "sqli",
					Cnt: 2,
				},
				LogLine{
					Ts: 9,
					Typ: "sqli",
					Cnt: 2,
				},
			},
			ts: 4,
			flag: 2,
			expected: 1,
		},
	}

	for _, test := range tests {
		actual := binarySearch(test.lines, test.ts, test.flag)
		if actual != test.expected {
			t.Errorf("expected: %d, but has %d\n", test.expected, actual)
		}
	}
}
