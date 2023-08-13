package parser

import (
	"errors"
	"fmt"
	"mileage-calculator/model"
	"mileage-calculator/time"
	"mileage-calculator/utils"
	"regexp"
	"strings"
)

func InputGetter(dataArr []model.CabData, line string) ([]model.CabData, error) {
	timeUnique := make(map[string]bool)

	// check if line is in correct format
	err := checkInputFormat(line)
	if err != nil {
		return dataArr, err
	}

	err = checkTimeLatest(dataArr, line)

	if err != nil {
		return dataArr, err
	}

	// if line is empty, break the loop
	if len(strings.TrimSpace(line)) == 0 {
		return dataArr, errors.New("error: empty line")
	}

	err = compareMileage(dataArr, line)
	lineSplit := strings.Split(line, " ")

	if ok := timeUnique[lineSplit[0]]; ok && len(dataArr) > 0 {
		fmt.Println("error: duplicate time")
		return dataArr, errors.New("error: duplicate time")
	} else {
		timeUnique[lineSplit[0]] = true
	}

	data := model.CabData{
		Time:    lineSplit[0],
		Mileage: lineSplit[1],
	}

	dataArr = append(dataArr, data)

	return dataArr, nil
}

func checkTimeLatest(dataArr []model.CabData, line string) error {
	if len(dataArr) > 0 && len(strings.TrimSpace(line)) > 0 {
		td := dataArr[len(dataArr)-1].Time
		ld := strings.Split(line, " ")
		timeDif, _ := time.TimeDiff(td, ld[0])
		res := time.TimeToInt(timeDif)

		//check if res is more than 5 minutes in milliseconds
		if res > 300000 {
			return errors.New("error: time should be within 5 minutes")
		}

		//check if res is negative
		if res <= 0 {
			return errors.New("error: time should be in ascending order")
		}

	}
	return nil
}

func checkInputFormat(line string) error {
	var regex, _ = regexp.Compile(`[0-9][0-9]:[0-9][0-9]:[0-9][0-9].[0-9]+ [0-9]+`)

	var isMatch = regex.MatchString(line)

	if !isMatch {
		return errors.New("error: invalid input")
	}
	return nil
}

func compareMileage(arr []model.CabData, mileage2 string) error {
	if len(arr) == 0 {
		return nil
	}

	mileageFloat1 := utils.StrToFloat(arr[len(arr)-1].Mileage)
	mileageFloat2 := utils.StrToFloat(mileage2)

	if mileageFloat2 < 0 {
		return errors.New("error: mileage should be positive")
	}

	if mileageFloat1 > mileageFloat2 {
		return errors.New("error: mileage should be in ascending order")
	}
	return nil
}
