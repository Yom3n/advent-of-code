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
	diff := getFirstItemsDiff(report)
	if diff < 0 {
		return isDscReportSafe(report)
	}
	if diff > 0 {
		return isAscReportSafe(report)
	}
	return false
}

// Checks first 2 or 3 items and based on them count difference.
// 3rd item is needed when first two are equal
func getFirstItemsDiff(report string) int {
	digits := strings.Fields(report)
	firstDigit := strToInt(digits[0])
	secondDigit := strToInt(digits[1])
	diff := secondDigit - firstDigit
	if diff == 0 {
		/// When no difference, we can ommit 1 record, because of ProblemDempener that accepts 1 error
		secondDigit = strToInt(digits[2])
		diff = secondDigit - firstDigit
	}
	return diff
}

func isDscReportSafe(r string) bool {
	return _validateReport(r, func(diff int) bool {
		fmt.Println(diff)
		return diff < 0 && -diff <= maxSaftyChange
	})
}

func isAscReportSafe(r string) bool {
	return _validateReport(r, func(diff int) bool {
		return diff > 0 && diff <= maxSaftyChange
	})
}

func _validateReport(r string, validCondition func(diff int) bool) bool {
	canSkipErrror := true
	digitStr := strings.Fields(r)
	latestDigit := strToInt(digitStr[0])
	for i := 1; i < len(digitStr); i++ {
		curr := strToInt(digitStr[i])
		diff := curr - latestDigit
		if validCondition(diff) {
			latestDigit = curr
			continue
		} else {
			if canSkipErrror {
				canSkipErrror = false
				if i == 1 {
					// Check if perhaps first digit should be ommited
					// if second and third are correct and first and second are wrong
					altDiff := strToInt(digitStr[2]) - curr
					if validCondition(altDiff) {
						// Skips first digit
						latestDigit = curr
					}
				}
				continue
			}
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
