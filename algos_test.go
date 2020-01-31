package main

import "testing"

func TestBinarySearch(t *testing.T) {
	A := []int{1, 4, 10, 20, 99}
	cases := []struct {
		input, expected int
	}{
		{4, 1},
		{99, 4},
		{87, -1},
	}
	for _, c := range cases {
		actual := BinarySearch(A, c.input)
		if actual != c.expected {
			t.Errorf("BinarySearch(%d) on %d == %d, expected %d", c.input, A, actual, c.expected)
		}
	}
}
