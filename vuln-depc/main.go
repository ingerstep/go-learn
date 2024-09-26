package main

import (
	"fmt"
	"os"

	"golang.org/x/text/language"
)

func main() {
	for _, arg := range os.Args[1:] {
		tag, err := language.Parse(arg)
		if err != nil {
			fmt.Printf("%s: error: %v\n", arg, err)
		} else if tag == language.Und {
			fmt.Printf("%s: undefined\n", arg)
		} else {
			fmt.Printf("%s: tag %s\n", arg, tag)
		}
	}
}
