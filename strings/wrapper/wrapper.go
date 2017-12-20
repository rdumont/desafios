package wrapper

import (
	"fmt"
	"math"
	"strings"
)

func Wrap(text string, limit int, justify bool) string {
	result := ""
	for _, p := range strings.Split(text, "\n") {
		lineLength := 0
		line := []string{}
		for _, word := range strings.Split(p, " ") {
			if lineLength+len(line)+len(word) > limit {
				result += printLine(line, limit, lineLength, justify)
				line = []string{word}
				lineLength = len(word)
				continue
			}

			line = append(line, word)
			lineLength += len(word)
		}

		result += printLine(line, limit, lineLength, justify)
	}

	return result
}

func printLine(line []string, limit, length int, justify bool) string {
	if !justify {
		return fmt.Sprintln(strings.Join(line, " "))
	}

	if length == 0 {
		return fmt.Sprintln()
	}

	if len(line) == 1 {
		return fmt.Sprintln(line[0])
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
