package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func InputGetter() []string {
	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)

	var lines []string
	//var timeSet  map[string]struct{}{}

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

		if len(lines) > 0 && len(strings.TrimSpace(line)) > 0 {
			td := strings.Split(lines[len(lines)-1], " ")
			ld := strings.Split(line, " ")
			timeDif, _ := TimeDiff(td[0], ld[0])
			res := TimeToInt(timeDif)

			//check if res is more than 5 minutes
			if res > 300000 {
				fmt.Println("error: more than 5 minutes")
				break
			}

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
