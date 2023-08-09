package main

import (
	"fmt"
	"strings"
)

func main() {
	lines := InputGetter()

	if len(lines) < 2 {
		fmt.Println("error: less than 2 lines")
		return
	}

	var t []string

	for i := 0; i < len(lines); i++ {
		s := strings.Split(lines[i], " ")
		t = append(t, s[0])
	}

	//compare time in t[i] and t[i+1] for i = 0 to len(t)
	var res []string
	for i := 0; i < len(t)-1; i++ {
		timeDiff, err := TimeDiff(t[i], t[i+1])
		if err != nil {
			fmt.Println(err)
			break
		}

		if err == nil {
			res = append(res, timeDiff)
		}
	}

	l := len(lines)
	lastMileage := strings.Split(lines[l-1], " ")
	//cast lastMileage[1] from string to float64
	lastMileageFloat := StrToFloat(lastMileage[1])

	fare := Fare(lastMileageFloat)

	fmt.Println(fare)
	fmt.Println(res)
}
