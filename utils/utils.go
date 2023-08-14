package utils

import (
	"fmt"
	"mileage-calculator/model"
)

func StrToFloat(str string) float64 {
	var float float64
	fmt.Sscanf(str, "%f", &float)

	return float
}

// FloatToStr converts float64 to string
func FloatToStr(float float64) string {
	return fmt.Sprintf("%f", float)
}

// PrintArrReversed prints the array of CabData in reverse order
func PrintArrReversed(arr []model.CabData) {
	for i := len(arr) - 1; i >= 0; i-- {
		price := FloatToStr(arr[i].Price)
		fmt.Println(arr[i].Time + " " + arr[i].Mileage + " " + price)
	}
}
