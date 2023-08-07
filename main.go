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

	res := TimeDiff(t[0], t[1])

	fmt.Println("output:")
	fmt.Println(lines)
	fmt.Println(res)
}
