package main

import (
	"fmt"
	"strconv"
	"strings"
)

// / Takes string containing 2 columns of ids that must be compared in format
// / XXXX XXXX
func compareLocationIds(locationIdsStr string) {
	locationIds := strings.Split(locationIdsStr, " ")
	notEmptyLocationIds := []int{}
	for _, idStr := range locationIds {
		if len(idStr) == 0 {
			fmt.Println("Empty id. Skipping")
			continue
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			// TODO its can't parse id for some reason. maybe because of /n? 
			fmt.Println(idStr)
			panic("Wrong id provided")
		}
		notEmptyLocationIds = append(notEmptyLocationIds, id)

	}
	list1 := []int{}
	list2 := []int{}
	for i, id := range notEmptyLocationIds {
		if i%2 == 0 {
			list1 = append(list1, id)
		} else {
			list2 = append(list2, id)
		}
		fmt.Println(list1)

		fmt.Println(list2)
	}
}
