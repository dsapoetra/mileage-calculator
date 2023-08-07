package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func InputGetter() []string {
	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)

	var lines []string
	for {
		// read line from stdin using newline as separator
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// if line is empty, break the loop
		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		//append the line to a slice
		lines = append(lines, line)
	}
	return lines
}

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

	time_and_ms := strings.Split(time, ".")

	time_split := strings.Split(time_and_ms[0], ":")

	//convert the hours and minutes to int
	hours := StrToInt(time_split[0])
	minutes := StrToInt(time_split[1])
	seconds := StrToInt(time_split[2])
	milisecond := StrToInt(time_and_ms[1])

	//convert the hours to milliseconds
	hours_to_milliseconds := hours * 3600000

	// convert the minutes to milliseconds
	minutes_to_milliseconds := minutes * 60000

	//convert the seconds to milliseconds
	seconds_to_milliseconds := seconds * 1000

	////add the hours and minutes
	total_minutes := hours_to_milliseconds + minutes_to_milliseconds + seconds_to_milliseconds + milisecond

	return total_minutes
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
	hours_str := fmt.Sprintf("%d", hours)
	minutes_str := fmt.Sprintf("%d", minutes)
	seconds_str := fmt.Sprintf("%d", seconds)
	miliseconds_str := fmt.Sprintf("%d", miliseconds)

	//add a zero to the hours and minutes if they are less than 10
	if hours < 10 {
		hours_str = fmt.Sprintf("0%d", hours)
	}
	if minutes < 10 {
		minutes_str = fmt.Sprintf("0%d", minutes)
	}

	//concatenate the hours and minutes to get the time string
	time_str := hours_str + ":" + minutes_str + ":" + seconds_str + "." + miliseconds_str + "000"

	return time_str
}

// a function that get time difference between two strings
func TimeDiff(time1 string, time2 string) string {
	//convert the time strings to int
	time1_int := TimeToInt(time1)
	time2_int := TimeToInt(time2)

	//get the difference between the two times
	diff := time2_int - time1_int

	//convert the difference to string
	diff_str := IntToTime(diff)

	return diff_str
}
