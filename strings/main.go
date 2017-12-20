package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
			if lineLength+len(line)+len(word) > limit {
				printLine(line, limit, lineLength)
				line = []string{word}
				lineLength = len(word)
				continue
			}

			line = append(line, word)
			lineLength += len(word)
		}

		printLine(line, limit, lineLength)
	}
}

func printLine(line []string, limit, length int) {
	if length == 0 {
		fmt.Println()
		return
	}

	spaceCount := limit - length
	wordSpaceCount := int(math.Trunc(float64(spaceCount / (len(line) - 1))))
	spaces := strings.Repeat(" ", wordSpaceCount)

	rem := spaceCount % (len(line) - 1)
	acc := ""
	for _, w := range line {
		if acc == "" {
			acc = w
			continue
		}

		if rem > 0 {
			acc += " "
			rem--
		}

		acc += spaces + w
	}

	fmt.Println(acc)
}
