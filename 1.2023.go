package main

import (
	"strconv"
	"strings"
)

// The newly-improved calibration document consists of lines of text;
// each line originally contained a specific calibration value that the Elves now need to recover.
// On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

// For example:

// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet

// In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.
// Consider your entire calibration document. What is the sum of all of the calibration values?

func CalibrateTrebuchet(calibrationDocument string) int {
	calibrationValues := []int{}
	lines := strings.Split(calibrationDocument, "\n")
	for _, line := range lines {
		calibrationVal := getLineCalibrationCode(line)
		calibrationValues = append(calibrationValues, calibrationVal)
	}
	return sum(calibrationValues)
}

func findFirstDigit(line string) int {
	for _, l := range line {
		digit, err := strconv.Atoi(string(l))
		if err != nil {
			continue
		}
		return digit
	}
	panic("Line should always have a digit")
}

func findLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		l := line[i]
		digit, err := strconv.Atoi(string(l))
		if err != nil {
			continue
		}
		return digit
	}
	panic("Line should always have a digit")
}

func getLineCalibrationCode(line string) int {
	firstDigit := findFirstDigit(line)
	lastDigit := findLastDigit(line)
	return firstDigit*10 + lastDigit
}

func sum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}
