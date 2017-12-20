package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wrap <limit> <file>")
		os.Exit(1)
	}

	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		panic(err)
	}

	for _, p := range strings.Split(string(bytes), "\n") {
		lineLength := 0
		line := []string{}
		for _, word := range strings.Split(p, " ") {
			if lineLength+len(word)+1 > limit {
				printLine(line)
				line = []string{word}
				lineLength = len(word)
				continue
			}

			line = append(line, word)
			lineLength += len(word) + 1
		}

		printLine(line)
	}
}

func printLine(line []string) {
	fmt.Println(strings.Join(line, " "))
}
