package wrapper

import (
	"fmt"
	"math"
	"strings"
)

func Wrap(text string, limit int) string {
	result := ""
	for _, p := range strings.Split(text, "\n") {
		lineLength := 0
		line := []string{}
		for _, word := range strings.Split(p, " ") {
			if lineLength+len(line)+len(word) > limit {
				result += printLine(line, limit, lineLength)
				line = []string{word}
				lineLength = len(word)
				continue
			}

			line = append(line, word)
			lineLength += len(word)
		}

		result += printLine(line, limit, lineLength)
	}

	return result
}

func printLine(line []string, limit, length int) string {
	if length == 0 {
		return fmt.Sprintln()
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

	return fmt.Sprintln(acc)
}
