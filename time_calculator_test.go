package main

import (
	"fmt"
	"testing"
)

// test for TimeToInt
func TestTimeToInt(t *testing.T) {
	//test for TimeToInt
	testTimeToInt := TimeToInt("00:00:00.000")
	if testTimeToInt != 0 {
		t.Errorf("TimeToInt was incorrect, got: %d, want: %d.", testTimeToInt, 0)
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

// test for TimeDiff
func TestTimeDiff(t *testing.T) {
	//test for TimeDiff
	testTimeDiff, _ := TimeDiff("00:00:00.0", "00:00:00.0")
	if testTimeDiff != "00:00:00.0" {
		t.Errorf("TimeDiff was incorrect, got: %s, want: %s.", testTimeDiff, "00:00:00.0")
	}
}
