package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// / Takes string containing 2 columns of ids that must be compared in format
// / XXXX XXXX
func CompareLocationIds(locationIdsStr string) {
	left, right := convertInputToLists(locationIdsStr)
	sortAsc(left)
	sortAsc(right)
	sum := 0
	for i, l := range left {
		p := right[i]
		sum += int(abs(l - p))
	}
	fmt.Println(sum)
}

// / Convers string input, and returns one list for the left column, and one for the right column
func convertInputToLists(locationIdsStr string) (left []int, right []int) {

	idsByLines := strings.Split(locationIdsStr, "\n")

	list1 := []int{}
	list2 := []int{}
	for _, idLine := range idsByLines {
		lineIds := strings.Split(idLine, " ")
		notEmptyLocationIds := []int{}
		for _, idStr := range lineIds {
			if len(idStr) == 0 {
				continue
			}
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println(idStr)
				panic("Wrong id provided")
			}
			notEmptyLocationIds = append(notEmptyLocationIds, id)

		}
		for i, id := range notEmptyLocationIds {
			if i%2 == 0 {
				list1 = append(list1, id)
			} else {
				list2 = append(list2, id)
			}
		}
	}
	return list1, list2
}

func sortAsc(list []int) {
	slices.SortFunc(list, func(i, j int) int {
		if i == j {
			return 0
		}
		if i < j {
			return -1
		}
		return 1
	})
}

// / Return absolute of the number
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
