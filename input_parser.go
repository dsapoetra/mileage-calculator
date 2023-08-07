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
