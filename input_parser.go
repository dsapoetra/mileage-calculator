package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func InputGetter(dataArr []Data, line string) ([]Data, error) {
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

	if len(dataArr) > 1 {
		err = compareMileage(dataArr[len(dataArr)-1].mileage, line)
	}
	lineSplit := strings.Split(line, " ")

	if ok := timeUnique[lineSplit[0]]; ok && len(dataArr) > 0 {
		fmt.Println("error: duplicate time")
		return dataArr, errors.New("error: duplicate time")
	} else {
		timeUnique[lineSplit[0]] = true
	}

	data := Data{
		time:    lineSplit[0],
		mileage: lineSplit[1],
	}

	dataArr = append(dataArr, data)

	return dataArr, nil
}

func checkTimeLatest(dataArr []Data, line string) error {
	if len(dataArr) > 0 && len(strings.TrimSpace(line)) > 0 {
		fmt.Println("HEREEE")
		td := dataArr[len(dataArr)-1].time
		fmt.Println("HEREEEE 2")
		ld := strings.Split(line, " ")
		timeDif, _ := TimeDiff(td, ld[0])
		res := TimeToInt(timeDif)
		fmt.Println(timeDif)
		fmt.Println(res)

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

func compareMileage(mileage1 string, mileage2 string) error {
	mileageFloat1 := StrToFloat(mileage1)
	mileageFloat2 := StrToFloat(mileage2)

	if mileageFloat2 < 0 {
		return errors.New("error: mileage should be positive")
	}

	if mileageFloat1 > mileageFloat2 {
		return errors.New("error: mileage should be in ascending order")
	}
	return nil
}
