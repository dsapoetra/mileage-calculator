package main

import (
	"fmt"
)

func main() {
	var dataArr []Data

	dataArr = InputGetter(dataArr)

	if len(dataArr) < 2 {
		fmt.Println("error: less than 2 lines")
		return
	}

	for i := 0; i < len(dataArr); i++ {
		milageFloat := StrToFloat(dataArr[i].mileage)
		dataArr[i].price = Fare(milageFloat)
	}
	PrintArrReversed(dataArr)
}
