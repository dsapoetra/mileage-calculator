package main

import (
	"bufio"
	"fmt"
	"mileage-calculator/fare_calculator"
	"mileage-calculator/model"
	"mileage-calculator/parser"
	"mileage-calculator/utils"
	"os"
)

func main() {
	dataArr := make([]model.CabData, 0)
	reader := bufio.NewReader(os.Stdin)
	var err error
	fmt.Println("input text:")
	for {
		line, _ := reader.ReadString('\n')
		dataArr, err = parser.InputGetter(dataArr, line)
		if err != nil {
			break
		}
	}

	if len(dataArr) < 2 {
		fmt.Println("error: less than 2 lines")
		return
	}

	for i := 0; i < len(dataArr); i++ {
		milageFloat := utils.StrToFloat(dataArr[i].Mileage)
		dataArr[i].Price = fare_calculator.Fare(milageFloat)
	}
	utils.PrintArrReversed(dataArr)
}
