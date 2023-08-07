package main

import (
	"fmt"
	"strings"
)

// a function that convert string to int
func StrToInt(str string) int {
	//convert the string to int
	var int int
	fmt.Sscanf(str, "%d", &int)

	return int
}

// a function that convert time string to int
func TimeToInt(time string) int {
	//split the string to get the hours and minutes

	timeAndMs := strings.Split(time, ".")

	timeSplit := strings.Split(timeAndMs[0], ":")

	//convert the hours and minutes to int
	hours := StrToInt(timeSplit[0])
	minutes := StrToInt(timeSplit[1])
	seconds := StrToInt(timeSplit[2])
	milisecond := StrToInt(timeAndMs[1])

	//convert the hours to milliseconds
	hoursToMilliseconds := hours * 3600000

	// convert the minutes to milliseconds
	minutesToMilliseconds := minutes * 60000

	//convert the seconds to milliseconds
	secondsToMilliseconds := seconds * 1000

	////add the hours and minutes
	totalMilliSeconds := hoursToMilliseconds + minutesToMilliseconds + secondsToMilliseconds + milisecond

	return totalMilliSeconds
}

// a function that convert int to time string
func IntToTime(int int) string {
	////convert the int to time string
	//hours := int / 60
	//minutes := int % 60

	//convert the int to time string, from milliseconds to hours, minutes, seconds and miliseconds
	hours := int / 3600000
	minutes := (int % 3600000) / 60000
	seconds := ((int % 3600000) % 60000) / 1000
	miliseconds := ((int % 3600000) % 60000) % 1000

	//convert the hours and minutes to string
	hoursStr := fmt.Sprintf("%d", hours)
	minutesStr := fmt.Sprintf("%d", minutes)
	secondsStr := fmt.Sprintf("%d", seconds)
	milisecondsStr := fmt.Sprintf("%d", miliseconds)

	//add a zero to the hours and minutes if they are less than 10
	if hours < 10 {
		hoursStr = fmt.Sprintf("0%d", hours)
	}
	if minutes < 10 {
		minutesStr = fmt.Sprintf("0%d", minutes)
	}

	if seconds < 10 {
		secondsStr = fmt.Sprintf("0%d", seconds)
	}

	//concatenate the hours and minutes to get the time string
	timeStr := hoursStr + ":" + minutesStr + ":" + secondsStr + "." + milisecondsStr + "000"

	return timeStr
}

// a function that get time difference between two strings
func TimeDiff(time1 string, time2 string) (string, error) {
	//convert the time strings to int
	time1_int := TimeToInt(time1)
	time2_int := TimeToInt(time2)

	//get the difference between the two times
	diff := time2_int - time1_int

	//check if diff is more than 5 minutes
	if diff > 300000 {
		return "", fmt.Errorf("The difference between the two times is more than 5 minutes")
	}

	//convert the difference to string
	diff_str := IntToTime(diff)

	return diff_str, nil
}
