package solution

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{[]int{-5}, []int{25}},
		{[]int{}, []int{}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{-2, -1}, []int{1, 4}},
	}

	for _, test := range tests {
		got := sortedSquares(test.input)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("sortedSquares(%v) = %v; expected %v", test.input, got, test.expected)
		}
	}
}
