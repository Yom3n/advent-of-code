package report_checker

import "testing"

func TestSecond2024(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	actual := GetNumSafeReports(input)
	expected := 4
	if actual != expected {
		t.Errorf("Should be: %v, Got: %v", expected, actual)
	}
}

func TestReportDsc(t *testing.T) {
	input := "7 6 4 2 1"
	actual := IsReportSafe(input)
	expected := true
	if actual != expected {
		t.Errorf("Expected: %v, Got: %v", expected, actual)
	}
}

func TestReportAsc(t *testing.T) {
	input := "1 3 6 7 9"
	actual := IsReportSafe(input)
	expected := true
	if actual != expected {
		t.Errorf("Expected: %v, Got: %v", expected, actual)
	}
}

/// With problem dampener you can ommit 1 bad digid in report
func TestProblemDampener(t *testing.T) {
	isSafe := IsReportSafe("1 3 2 4 5")
	if isSafe != true {
		t.Errorf("1 3 2 4 5 is safe report")
	}
}

func TestZeroDiff(t *testing.T) {
	isSafe := IsReportSafe("1 1 2 4 5")
	if isSafe != true {
		t.Errorf("1 3 2 4 5 is safe report")
	}
}
