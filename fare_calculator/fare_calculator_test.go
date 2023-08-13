package fare_calculator

import (
	"math"
	"testing"
)

func TestFare(t *testing.T) {
	epsilon := 1e-5
	//test for Fare
	testFare := Fare(1000)
	if math.Abs(testFare-400.0) > epsilon {
		t.Errorf("Fare was incorrect, got: %v, want: %f.", testFare, 400.0)
	}

	testFare = Fare(1500)
	if math.Abs(testFare-450.0) > epsilon {
		t.Errorf("Fare was incorrect, got: %v, want: %f.", testFare, 450.0)
	}
	//test for Fare
	testFare = Fare(10000)
	if math.Abs(testFare-1300.0) > epsilon {
		t.Errorf("Fare was incorrect, got: %v, want: %f.", testFare, 1300.0)
	}

	// test for Fare 10001
	testFare = Fare(15000)
	if math.Abs(testFare-1871.428571) > epsilon {
		t.Errorf("Fare was incorrect, got: %v, want: %f.", testFare, 1852.0)
	}
}
