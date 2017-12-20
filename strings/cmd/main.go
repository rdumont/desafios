package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/idwall/desafios/strings/wrapper"
)

func main() {
	fs := flag.NewFlagSet("bla", flag.ContinueOnError)
	fs.Usage = func() {}

	justify := fs.Bool("j", false, "-j")

	if err := fs.Parse(os.Args[1:]); err != nil {
		exit(1)
	}

	if fs.NArg() < 2 {
		fmt.Println("Error: wrong number of arguments")
		exit(1)
	}

	limit, err := strconv.Atoi(fs.Arg(0))
	if err != nil {
		fmt.Println("Error:", err)
		exit(1)
	}

	bytes, err := ioutil.ReadFile(fs.Arg(1))
	if err != nil {
		fmt.Println("Error:", err)
		exit(1)
	}

	result := wrapper.Wrap(string(bytes), limit, *justify)
	fmt.Print(result)
}

func exit(code int) {
	fmt.Println(`Usage: wrap [options] <limit> <file>

Options:
  -j	Justify text
`)

	os.Exit(code)
}
