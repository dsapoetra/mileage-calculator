package main

import (
	"fmt"
	"testing"
)

// test for StrToInt
func TestStrToInt(t *testing.T) {
	//test for StrToInt
	testStrToInt := StrToInt("10")
	if testStrToInt != 10 {
		t.Errorf("StrToInt was incorrect, got: %d, want: %d.", testStrToInt, 10)
	}
}

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
	if testIntToTime != "00:00:00.0000" {
		t.Errorf("IntToTime was incorrect, got: %s, want: %s.", testIntToTime, "00:00:00.0000")
	}
}

// test for TimeDiff
func TestTimeDiff(t *testing.T) {
	//test for TimeDiff
	testTimeDiff := TimeDiff("00:00:00.000", "00:00:00.000")
	if testTimeDiff != "00:00:00.0000" {
		t.Errorf("TimeDiff was incorrect, got: %s, want: %s.", testTimeDiff, "00:00:00.000")
	}
}
