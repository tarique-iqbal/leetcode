package solution

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
		{[]int{2, 3, 10, 5, 7, 8, 9}, 36},
	}

	for _, tt := range tests {
		result := maxArea(tt.height)
		if result != tt.expected {
			t.Errorf("MaxArea(%v) = %d; want %d", tt.height, result, tt.expected)
		}
	}
}
