package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dataArr := make([]Data, 0)
	reader := bufio.NewReader(os.Stdin)
	var err error
	fmt.Println("input text:")
	for {
		line, _ := reader.ReadString('\n')
		dataArr, err = InputGetter(dataArr, line)
		if err != nil {
			break
		}
	}

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
