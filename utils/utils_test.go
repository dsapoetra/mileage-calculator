package utils

import (
	"testing"
)

// test for StrToFloat
func TestStrToFloat(t *testing.T) {
	//test for StrToFloat
	testStrToFloat := StrToFloat("10.0")
	if testStrToFloat != 10.0 {
		t.Errorf("StrToFloat was incorrect, got: %f, want: %f.", testStrToFloat, 10.0)
	}
}

// test for FloatToStr
func TestFloatToStr(t *testing.T) {
	//test for FloatToStr
	testFloatToStr := FloatToStr(10.0)
	if testFloatToStr != "10.000000" {
		t.Errorf("FloatToStr was incorrect, got: %s, want: %s.", testFloatToStr, "10.000000")
	}
}
