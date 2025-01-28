package main

// Takes lesf ot integers
// Returns list of sorted integers using merge sort algorithm
// Efficiency: n*Log(n)
func MergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	left, right := splitList(input)
	left = MergeSort(left)
	right = MergeSort(right)
	return mergeLists(left, right)
}

func splitList(input []int) ([]int, []int) {
	midIndex := len(input) / 2
	return input[0:midIndex], input[midIndex:]
}

func mergeLists(left []int, right []int) []int {
	leftIndex := 0
	rightIndex := 0
	sorted := []int{}
	for {
		if left[leftIndex] <= right[rightIndex] {
			sorted = append(sorted, left[leftIndex])
			leftIndex++
		} else {
			sorted = append(sorted, right[rightIndex])
			rightIndex++
		}
		if leftIndex >= len(left) && rightIndex >= len(right) {
			return sorted
		}
		if leftIndex >= len(left) && rightIndex < len(right) {
			sorted = append(sorted, right[rightIndex:]...)
			return sorted
		}
		if rightIndex >= len(right) && leftIndex < len(left) {
			sorted = append(sorted, left[leftIndex:]...)
			return sorted
		}

	}

}
