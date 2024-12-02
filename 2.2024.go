package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetNumSafeReports(reports string) int {
	reportsPerLine := strings.Split(reports, "\n")
	numSafe := 0
	for _, report := range reportsPerLine {
		if isReportSafe(report) {
			numSafe++
		}
	}
	return numSafe
}

const maxSaftyChange = 3

// / Check single report
func isReportSafe(report string) bool {
	digitStr := strings.Fields(report)
	lastDigit := strToInt(digitStr[0])
	for i := 1; i < len(digitStr); i++ {
		curr := strToInt(digitStr[i])
		diff := lastDigit - curr
		if diff > 0 && diff <= maxSaftyChange {
			lastDigit = curr
			break
		} else {
			return false
		}
	}
	return true
}

func strToInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(str)
		panic("Wrong value")
	}
	return v
}
