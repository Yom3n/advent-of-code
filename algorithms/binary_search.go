package main

// Gets sorted list of values, and target item
// Returns index of target, or -1 when target not found
// Complexity Log(n)
func BinarySearch(values []int, target int) (index int) {
	minIndex := 0
	maxIndex := len(values) - 1
	for minIndex <= maxIndex {
		midIndex := minIndex + ((maxIndex - minIndex) / 2)
		midValue := values[midIndex]
		if target > midValue {
			minIndex = midIndex + 1
			continue
		} else if target < midValue {
			maxIndex = midIndex - 1
			continue

		}
		return midIndex
	}
	return -1
}
