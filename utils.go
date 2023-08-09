package main

import "fmt"

// a function that convert string to int
func StrToInt(str string) int {
	//convert the string to int
	var int int
	fmt.Sscanf(str, "%d", &int)

	return int
}

func StrToFloat(str string) float64 {
	var float float64
	fmt.Sscanf(str, "%f", &float)

	return float
}
