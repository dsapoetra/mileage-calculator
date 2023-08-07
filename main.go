package main

import (
	"fmt"
	"strings"
)

func main() {
	lines := InputGetter()
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

	fmt.Println("output:")
	fmt.Println(lines)
	fmt.Println(res)
}
