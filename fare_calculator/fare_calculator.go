package fare_calculator

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// Fare a function to calculate fare_calculator with input mileage
// Input: mileage in float64 as meter
// Output: fare_calculator in float64
// 1. The base Fare is 400 yen for up to 1 km.
// 2. Up to 10 km, 40 yen is added every 400 meters.
// 3. Over 10km, 40 yen is added every 350 meters.
func Fare(mileage float64) decimal.Decimal {
	mile := decimal.NewFromFloat(mileage)

	if mile.LessThanOrEqual(decimal.NewFromInt(1000)) {
		return decimal.NewFromInt(400)
	} else if mile.LessThanOrEqual(decimal.NewFromInt(10000)) {
		//fare = 400 + ((mileage-1000)/400)*40
		base := decimal.NewFromInt(400)
		mileageMinusThousand := mile.Sub(decimal.NewFromInt(1000))
		mileageMinusThousandDivFourHundred := mileageMinusThousand.Div(decimal.NewFromInt(400))
		mileageMinusThousandDivFourHundredTimeForty := mileageMinusThousandDivFourHundred.Mul(decimal.NewFromInt(40))
		total := base.Add(mileageMinusThousandDivFourHundredTimeForty)

		fmt.Println(total)

		return total

	} else {
		//mileage 15000
		//fare := 400 + (9000/400)*40 + ((mileage-10000)/350)*40
		b := 400.0
		per400 := 9000.0 / 400.0
		t40 := per400 * 40
		min10000 := mileage - 10000
		per350 := min10000 / 350
		time40 := per350 * 40
		r := b + t40 + time40
		fmt.Println("1: ", r, "mileage: ", mileage)
		//fare := 400 + (9000/400)*40 + ((mileage-10000)/350)*40
		b = 400.0
		per400 = 9000.0 / 400.0
		t40 = per400 * 40
		min10000 = 15000 - 10000
		per350 = min10000 / 350
		time40 = per350 * 40
		r = b + t40 + time40
		fmt.Println("2: ", r)
		//fare = 400 + (9000/400)*40 + (5000/350)*40
		//fmt.Println("3: ", fare)

		base := decimal.NewFromInt(400)
		NineThousandPerFourHundred := decimal.NewFromInt(9000).Div(decimal.NewFromInt(400))
		NineThousandPerFourHundredTimesForty := NineThousandPerFourHundred.Mul(decimal.NewFromInt(40))
		mileageMinusTenThousand := mile.Sub(decimal.NewFromInt(10000))
		mileageMinusTenThousandDivThreeFifty := mileageMinusTenThousand.Div(decimal.NewFromInt(350))
		mileageMinusTenThousandDivFourHundredTimeForty := mileageMinusTenThousandDivThreeFifty.Mul(decimal.NewFromInt(40))
		total := base.Add(NineThousandPerFourHundredTimesForty).Add(mileageMinusTenThousandDivFourHundredTimeForty)

		fmt.Println(total)
		return total
	}
}
