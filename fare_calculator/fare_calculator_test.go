package fare_calculator

import (
	"fmt"
	"testing"
)

// test for Fare
func TestFare(t *testing.T) {
	//test for Fare
	testFare := Fare(1000)
	if testFare != 400 {
		t.Errorf("Fare was incorrect, got: %f, want: %f.", testFare, 400.0)
	}
	//test for Fare
	testFare = Fare(1001)
	if testFare != 400.1 {
		t.Errorf("Fare was incorrect, got: %f, want: %f.", testFare, 400.1)
	}
	//test for Fare
	testFare = Fare(10000)
	if testFare != 1300.0 {
		t.Errorf("Fare was incorrect, got: %f, want: %f.", testFare, 1300.0)
	}

	// test for Fare 10001
	testFare = Fare(10001)
	fmt.Println(testFare)
	if testFare != 1400.1142857142856 {
		t.Errorf("Fare was incorrect, got: %f, want: %f.", testFare, 1400.1142857142856)
	}
}
