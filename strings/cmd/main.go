package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/idwall/desafios/strings/wrapper"
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

	result := wrapper.Wrap(string(bytes), limit)
	fmt.Print(result)
}
