package report_checker

import (
	"fmt"
	"strconv"
	"strings"
)

const MaxLevelDiff = 3
const AcceptableErrorsAmount = 1

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

// / Check single report
func IsReportSafe(report string) bool {
	if isDscReportSafe(report) {
		return true
	}
	return isAscReportSafe(report)
}

func isDscReportSafe(r string) bool {
	numErrors := _getReportNumErrors(r, func(diff int) bool {
		return diff < 0 && -diff <= MaxLevelDiff
	})
	fmt.Println("Dsc errors")
	fmt.Println(numErrors)
	return numErrors <= AcceptableErrorsAmount
}

func isAscReportSafe(r string) bool {
	numErrors := _getReportNumErrors(r, func(diff int) bool {
		return diff > 0 && diff <= MaxLevelDiff
	})
	fmt.Println("Asc errors")
	fmt.Println(numErrors)
	return numErrors <= AcceptableErrorsAmount
}

func _getReportNumErrors(r string, validCondition func(diff int) bool) int {
	errors := 0
	digitStr := strings.Fields(r)
	latestDigit := strToInt(digitStr[0])
	for i := 1; i < len(digitStr); i++ {
		curr := strToInt(digitStr[i])
		diff := curr - latestDigit
		if validCondition(diff) {
			latestDigit = curr
			continue
		} else {
			if i == 1 {
				// Checks if first item should be ommited.
				// This implementation won't work if [AcceptedErrorsAmount] will be increased
				altDif := strToInt(digitStr[2]) - curr
				if validCondition(altDif) {
					latestDigit = curr
				}
			}
			errors += 1
		}
	}
	return errors
}

func strToInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(str)
		panic("Wrong value")
	}
	return v
}
