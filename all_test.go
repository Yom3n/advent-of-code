package main

import "testing"

func TestSecond2024(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	actual := GetNumSafeReports(input)
	expected := 2
	if actual != expected {
		t.Errorf("Should be: %v, Got: %v", expected, actual)
	}
}

func TestReport(t *testing.T) {
	input := "9 7 6 2 1"
	actual := IsReportSafe(input)
	expected := false
	if actual != expected {
		t.Errorf("Expected: %v, Got: %v", expected, actual)
	}
}
