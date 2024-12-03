package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO Nie ogarnalem, ze wszysteki moga sie zwieszkac lub wszystkie moga sie zmniejszac
func GetNumSafeReports(reports string) int {
	reportsPerLine := strings.Split(reports, "\n")
	numSafe := 0
	for _, report := range reportsPerLine {
		if IsReportSafe(report) {
			fmt.Println("SAFE")
			numSafe++
			continue
		}
		fmt.Println("UNSAFE")
	}
	return numSafe
}

const maxSaftyChange = 3

// / Check single report
func IsReportSafe(report string) bool {
	diff := report[1] - report[0]
	if diff > 0 {
		return isDscReportSafe(report)
	}
	if diff < 0 {
		return isAscReportSafe(report)
	}
	return false
}

func isDscReportSafe(r string) bool {
	digitStr := strings.Fields(r)
	lastDigit := strToInt(digitStr[0])
	for i := 1; i < len(digitStr); i++ {
		curr := strToInt(digitStr[i])
		diff := lastDigit - curr
		fmt.Println(diff)
		if diff > 0 && diff <= maxSaftyChange {
			lastDigit = curr
			continue
		} else {
			return false
		}
	}
	return true
}

func isAscReportSafe(r string) bool {
	digitStr := strings.Fields(r)
	lastDigit := strToInt(digitStr[0])
	for i := 1; i < len(digitStr); i++ {
		curr := strToInt(digitStr[i])
		diff := curr - lastDigit
		if diff > 0 && diff <= maxSaftyChange {
			lastDigit = curr
			continue
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
