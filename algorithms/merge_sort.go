package main

// Takes lesf ot integers
// Returns list of sorted integers using merge sort algorithm
// Efficiency: n*Log(n)
func MergeSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	midIndex := len(input) / 2
	left := input[:midIndex]
	right := input[midIndex:]
	left = MergeSort(left)
	right = MergeSort(right)
	return Merge(left, right)

}

func Merge(left []int, right []int) []int {
	leftIndex := 0
	rightIndex := 0
	output := []int{}
	for leftIndex < len(left) && rightIndex < len(right) {
		leftValue := left[leftIndex]
		rightValue := right[rightIndex]
		if leftValue <= rightValue {
			output = append(output, leftValue)
			leftIndex++
			continue
		} else {
			output = append(output, rightValue)
			rightIndex++
			continue
		}
	}
	for leftIndex < len(left) {
		output = append(output, left[leftIndex])
		leftIndex++
	}
	for rightIndex < len(right) {
		output = append(output, right[rightIndex])
		rightIndex++
	}
	return output
}
