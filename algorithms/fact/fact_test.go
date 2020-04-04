package fact

import "testing"

func TestFact(t *testing.T) {
	tests := []struct {
		i   int
		ret int
	}{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
	}

	for _, test := range tests {
		actual := Fact(test.i)
		if test.ret != actual {
			t.Errorf("Fact(%d) = %d", test.i, actual)
		}
	}
}
