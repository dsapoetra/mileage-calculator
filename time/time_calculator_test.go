package time

import (
	"fmt"
	"testing"
)

// test for ToInt
func TestTimeToInt(t *testing.T) {
	//test for ToInt
	testTimeToInt := ToInt("00:00:00.000")
	if testTimeToInt != 0 {
		t.Errorf("ToInt was incorrect, got: %d, want: %d.", testTimeToInt, 0)
	}
}

// test for IntToTime
func TestIntToTime(t *testing.T) {
	//test for IntToTime
	testIntToTime := IntToTime(0)
	fmt.Println(testIntToTime)
	if testIntToTime != "00:00:00.0" {
		t.Errorf("IntToTime was incorrect, got: %s, want: %s.", testIntToTime, "00:00:00.0")
	}
}

// test for Diff
func TestTimeDiff(t *testing.T) {
	//test for Diff
	testTimeDiff, _ := Diff("00:00:00.0", "00:00:00.0")
	if testTimeDiff != "00:00:00.0" {
		t.Errorf("Diff was incorrect, got: %s, want: %s.", testTimeDiff, "00:00:00.0")
	}
}
