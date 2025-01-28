package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse order",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted array",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			expected: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Duplicate elements",
			input:    []int{5, 3, 3, 1, 4, 5},
			expected: []int{1, 3, 3, 4, 5, 5},
		},
		{
			name:     "Negative numbers",
			input:    []int{-3, -1, -2, -4, -5},
			expected: []int{-5, -4, -3, -2, -1},
		},
		{
			name:     "Mixed positive and negative numbers",
			input:    []int{-3, 10, 5, -1, 0, 2},
			expected: []int{-3, -1, 0, 2, 5, 10},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MergeSort(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("MergeSort(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}
