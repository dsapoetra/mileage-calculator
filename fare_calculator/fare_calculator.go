package fare_calculator

// a function to calculate fare_calculator with input mileage
// Input: mileage in float64 as meter
// Output: fare_calculator in float64
// 1. The base fare_calculator is 400 yen for up to 1 km.
// 2. Up to 10 km, 40 yen is added every 400 meters.
// 3. Over 10km, 40 yen is added every 350 meters.
func Fare(mileage float64) float64 {
	var fare float64
	if mileage <= 1000 {
		fare = 400
	} else if mileage <= 10000 {
		fare = 400 + ((mileage-1000)/400)*40
	} else {
		fare = 400 + (10000/400)*40 + ((mileage-10000)/350)*40
	}

	return fare
}
