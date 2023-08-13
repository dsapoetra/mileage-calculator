package time

import (
	"fmt"
	"strconv"
	"strings"
)

// TimeToInt converts time string to int as milliseconds
func TimeToInt(time string) int {

	// Split input from hours, minutes, seconds and milliseconds
	timeAndMs := strings.Split(time, ".")

	// Split input from timeAndMs to hours, minutes and seconds
	timeSplit := strings.Split(timeAndMs[0], ":")

	// convert the hours, minutes, seconds and milliseconds to int
	hours, _ := strconv.Atoi(timeSplit[0])
	minutes, _ := strconv.Atoi(timeSplit[1])
	seconds, _ := strconv.Atoi(timeSplit[2])
	millisecond, _ := strconv.Atoi(timeAndMs[1])

	//convert the hours to milliseconds
	hoursToMilliseconds := hours * 3600000

	// convert the minutes to milliseconds
	minutesToMilliseconds := minutes * 60000

	//convert the seconds to milliseconds
	secondsToMilliseconds := seconds * 1000

	////add the hours and minutes
	totalMilliSeconds := hoursToMilliseconds + minutesToMilliseconds + secondsToMilliseconds + millisecond

	return totalMilliSeconds
}

// IntToTime converts int as milliseconds to time string
func IntToTime(int int) string {

	//convert the int to time string, from milliseconds to hours, minutes, seconds and milliseconds
	hours := int / 3600000
	minutes := (int % 3600000) / 60000
	seconds := ((int % 3600000) % 60000) / 1000
	milliseconds := ((int % 3600000) % 60000) % 1000

	//convert the hours and minutes to string
	hoursStr := fmt.Sprintf("%d", hours)
	minutesStr := fmt.Sprintf("%d", minutes)
	secondsStr := fmt.Sprintf("%d", seconds)
	millisecondsStr := fmt.Sprintf("%d", milliseconds)

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
	timeStr := hoursStr + ":" + minutesStr + ":" + secondsStr + "." + millisecondsStr

	return timeStr
}

// TimeDiff gets the difference between two time strings
func TimeDiff(time1 string, time2 string) (string, error) {
	//convert the time strings to int
	time1Int := TimeToInt(time1)
	time2Int := TimeToInt(time2)

	//get the difference between the two times
	diff := time2Int - time1Int

	//convert the difference to string
	diffStr := IntToTime(diff)

	return diffStr, nil
}
