package solution

import "testing"

func TestNumRescueBoats(t *testing.T) {
	tests := []struct {
		name     string
		people   []int
		limit    int
		expected int
	}{
		{
			name:     "Example case 1",
			people:   []int{1, 2},
			limit:    3,
			expected: 1,
		},
		{
			name:     "Example case 2",
			people:   []int{3, 2, 2, 1},
			limit:    3,
			expected: 3,
		},
		{
			name:     "Example case 3",
			people:   []int{3, 5, 3, 4},
			limit:    5,
			expected: 4,
		},
		{
			name:     "All people under half limit",
			people:   []int{1, 1, 1, 1},
			limit:    10,
			expected: 2,
		},
		{
			name:     "Each person needs a boat",
			people:   []int{5, 5, 5, 5},
			limit:    5,
			expected: 4,
		},
		{
			name:     "Empty list",
			people:   []int{},
			limit:    10,
			expected: 0,
		},
		{
			name:     "Single person",
			people:   []int{2},
			limit:    3,
			expected: 1,
		},
		{
			name:     "All pairs fit",
			people:   []int{2, 2, 2, 2},
			limit:    4,
			expected: 2,
		},
		{
			name:     "One heavy person prevents pairing",
			people:   []int{1, 2, 3, 9},
			limit:    5,
			expected: 3,
		},
	}

	for _, tt := range tests {
		result := numRescueBoats(tt.people, tt.limit)
		if result != tt.expected {
			t.Errorf("Expected %d, got %d", tt.expected, result)
		}
	}
}
