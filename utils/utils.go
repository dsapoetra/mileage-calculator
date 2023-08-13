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

// float to string
func FloatToStr(float float64) string {
	return fmt.Sprintf("%f", float)
}

func PrintArrReversed(arr []model.CabData) {
	for i := len(arr) - 1; i >= 0; i-- {
		price := arr[i].Price.String()
		fmt.Println(arr[i].Time + " " + arr[i].Mileage + " " + price)
	}
}
