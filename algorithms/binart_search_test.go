package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name       string
		valuesList []int
		target     int
		expected   int
	}{
		{
			name:       "Univen input, target on the right",
			valuesList: []int{1, 43, 59, 88, 89, 100, 123, 432, 666},
			target:     432,
			expected:   7,
		},
		{
			name:       "Univen input, target on the left",
			valuesList: []int{1, 43, 59, 88, 89, 100, 123, 432, 666},
			target:     59,
			expected:   2,
		},
		{
			name:       "Univen input, target in the middle",
			valuesList: []int{1, 43, 59, 88, 89, 100, 123, 432, 666},
			target:     89,
			expected:   4,
		},
		{
			name:       "Univen input, target in first",
			valuesList: []int{1, 43, 59, 88, 89, 100, 123, 432, 666},
			target:     1,
			expected:   0,
		},
		{
			name:       "Univen input, target in target last",
			valuesList: []int{1, 43, 59, 88, 89, 100, 123, 432, 666},
			target:     666,
			expected:   8,
		},
		{
			name:       "Univen input, target not exists",
			valuesList: []int{1, 43, 59, 88, 89, 100, 123, 432, 666},
			target:     555,
			expected:   -1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := BinarySearch(test.valuesList, test.target)
			if actual != test.expected {
				t.Errorf("Binary search values: %v, target: %v, actual: %v, expected: %v",
					test.valuesList,
					test.target,
					actual,
					test.expected,
				)
			}
		})
	}

}
