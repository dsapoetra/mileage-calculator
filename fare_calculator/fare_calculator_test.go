package fare_calculator

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestFare(t *testing.T) {
	//test for Fare
	testFare := Fare(1000)
	if testFare != decimal.NewFromInt(400) {
		t.Errorf("Fare was incorrect, got: %v, want: %f.", testFare, 400.0)
	}

	testFare = Fare(1500)
	if testFare != decimal.NewFromFloat(450.0) {
		t.Errorf("Fare was incorrect, got: %v, want: %f.", testFare, 450.0)
	}
	//test for Fare
	testFare = Fare(10000)
	if testFare != decimal.NewFromInt(1300) {
		t.Errorf("Fare was incorrect, got: %v, want: %f.", testFare, 1300.0)
	}

	// test for Fare 10001
	testFare = Fare(15000)
	if testFare.Equal(decimal.NewFromFloat(1871.428571)) {
		t.Errorf("Fare was incorrect, got: %v, want: %f.", testFare, 1852.0)
	}
}
