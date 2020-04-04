package set

import (
	"reflect"
	"testing"
)

func TestIntersectionTwo(t *testing.T) {
	tests := []struct {
		s1, s2, expected Set
	}{
		{Set{}, Set{}, Set{}},
		{Set{1, 2}, Set{}, Set{}},
		{Set{1, 2}, Set{3, 4}, Set{}},
		{Set{1, 2, 3, 4}, Set{1, 2}, Set{1, 2}},
	}

	for _, test := range tests {
		s := Sort(IntersectionTwo(test.s1, test.s2))
		// fmt.Printf("%#v, %#v\n", s, test.expected)
		if !reflect.DeepEqual(test.expected, s) {
			t.Errorf("IntersectionTwo(%v, %v) = %v, but exepect %v",
				test.s1, test.s2, s, test.expected)
		}
	}

	// s1 := Set{1, 2, 3, 4}
	// s2 := Set{2, 3, 4, 5, 6}
	// s := IntersectionTwo(s1, s2)
	// if !reflect.DeepEqual(s, Set{2, 3, 4}) {
	// 	t.Errorf("IntersectionTwo(%v, %v) = %v, but exepect %v",
	// 		s1, s2, s, Set{2, 3, 4})
	// }
}

func TestIntersection(t *testing.T) {
	s1 := Set{1, 2}
	s2 := Set{1, 2, 3}
	s3 := Set{2, 3}
	s := Sort(Intersection(s1, s2, s3))
	if !reflect.DeepEqual(s, Set{2}) {
		t.Errorf("Intersection(%v, %v, %v) = %v, but expected %v",
			s1, s2, s3, s, Set{2})
	}
}

func TestUnionTwo(t *testing.T) {
	tests := []struct {
		s1, s2, expected Set
	}{
		{Set{}, Set{}, Set{}},
		{Set{1, 2}, Set{}, Set{1, 2}},
		{Set{1, 2}, Set{3, 4}, Set{1, 2, 3, 4}},
		{Set{5, 1, 2, 3, 4}, Set{1, 2}, Set{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		s := Sort(UnionTwo(test.s1, test.s2))
		// fmt.Printf("%#v, %#v\n", s, test.expected)
		if !reflect.DeepEqual(test.expected, s) {
			t.Errorf("UnionTwo(%v, %v) = %v, but exepect %v",
				test.s1, test.s2, s, test.expected)
		}
	}
}

func TestUnion(t *testing.T) {
	s1 := Set{1, 2}
	s2 := Set{1, 2, 3}
	s3 := Set{5, 4, 2, 3}
	s := Sort(Union(s1, s2, s3))
	if !reflect.DeepEqual(s, Set{1, 2, 3, 4, 5}) {
		t.Errorf("Union(%v, %v, %v) = %v, but expected %v",
			s1, s2, s3, s, Set{1, 2, 3, 4, 5})
	}
}

func TestDiffTwo(t *testing.T) {
	tests := []struct {
		s1, s2, expected Set
	}{
		{Set{}, Set{}, Set{}},
		{Set{1, 2}, Set{}, Set{1, 2}},
		{Set{1, 2}, Set{3, 4}, Set{1, 2, 3, 4}},
		{Set{5, 1, 2, 3, 4}, Set{1, 2}, Set{3, 4, 5}},
	}

	for _, test := range tests {
		s := Sort(DiffTwo(test.s1, test.s2))
		// fmt.Printf("%#v, %#v\n", s, test.expected)
		if !reflect.DeepEqual(test.expected, s) {
			t.Errorf("UnionTwo(%v, %v) = %v, but exepect %v",
				test.s1, test.s2, s, test.expected)
		}
	}
}

func TestDiff(t *testing.T) {
	s1 := Set{1, 2}
	s2 := Set{1, 2, 3}
	s3 := Set{5, 4, 2, 3}
	s := Sort(Union(s1, s2, s3))
	if !reflect.DeepEqual(s, Set{1, 2, 3, 4, 5}) {
		t.Errorf("Union(%v, %v, %v) = %v, but expected %v",
			s1, s2, s3, s, Set{1, 2, 3, 4, 5})
	}
}
