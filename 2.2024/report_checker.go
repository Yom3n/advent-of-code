package report_checker 

import (
	"fmt"
	"strconv"
	"strings"
)

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
	digits := strings.Fields(report)
	firstDigit := strToInt(digits[0])
	secondDigit := strToInt(digits[1])
	diff := secondDigit - firstDigit
	if diff < 0 {
		return isDscReportSafe(report)
	}
	if diff > 0 {
		return isAscReportSafe(report)
	}
	return false
}

func isDscReportSafe(r string) bool {
	digitStr := strings.Fields(r)
	lastDigit := strToInt(digitStr[0])
	for i := 1; i < len(digitStr); i++ {
		curr := strToInt(digitStr[i])
		diff := curr - lastDigit
		if diff < 0 && -diff <= maxSaftyChange {
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
