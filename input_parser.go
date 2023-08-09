package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func InputGetter(dataArr []Data) []Data {
	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)
	i := int64(0)

	for {
		// read line from stdin using newline as separator

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		var regex, _ = regexp.Compile(`[0-9][0-9]:[0-9][0-9]:[0-9][0-9].[0-9]+ [0-9]+`)

		var isMatch = regex.MatchString(line)

		if !isMatch {
			fmt.Println("error: invalid input")
			break
		}

		if len(dataArr) > 0 && len(strings.TrimSpace(line)) > 0 {
			td := dataArr[len(dataArr)-1].time
			ld := strings.Split(line, " ")
			timeDif, _ := TimeDiff(td, ld[0])
			res := TimeToInt(timeDif)

			//check if res is more than 5 minutes in milliseconds
			if res > 300000 {
				fmt.Println("error: more than 5 minutes")
				break
			}

			//check if res is negative
			if res < 0 {
				fmt.Println("error: time should be in ascending order")
				break
			}

		}

		// if line is empty, break the loop
		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		// compare mileage with previous line
		if len(dataArr) > 0 {
			mileage := StrToFloat(dataArr[len(dataArr)-1].mileage)
			lineSplit := strings.Split(line, " ")
			mileage2 := StrToFloat(lineSplit[1])

			//check if mileage is negative
			if mileage2 < 0 {
				fmt.Println("error: mileage should be positive")
				break
			}

			//check if mileage is in ascending order
			if mileage2 < mileage {
				fmt.Println("error: mileage should be in ascending order")
				break
			}
		}

		lineSplit := strings.Split(line, " ")
		data := Data{
			time:    lineSplit[0],
			mileage: lineSplit[1],
		}

		dataArr = append(dataArr, data)

		i++

	}
	return dataArr
}
