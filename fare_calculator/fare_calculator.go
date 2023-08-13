package fare_calculator

// Fare a function to calculate fare_calculator with input mileage
// Input: mileage in float64 as meter
// Output: fare_calculator in float64
// 1. The base Fare is 400 yen for up to 1 km.
// 2. Up to 10 km, 40 yen is added every 400 meters.
// 3. Over 10km, 40 yen is added every 350 meters.
func Fare(mileage float64) float64 {

	if mileage <= 1000 {
		return 400.0
	} else if mileage > 1000 && mileage <= 10000 {
		base := 400.0
		minusThousand := mileage - 1000
		per400 := minusThousand / 400.0
		t40 := per400 * 40
		total := base + t40

		return total

	} else {

		b := 400.0
		per400 := 9000.0 / 400.0
		t40 := per400 * 40
		min10000 := mileage - 10000
		per350 := min10000 / 350
		time40 := per350 * 40
		total := b + t40 + time40
		return total
	}
}
