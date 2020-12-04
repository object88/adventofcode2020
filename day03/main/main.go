package main

import (
	"fmt"
	"os"

	"github.com/object88/adventofcode2020/day03"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Expected: %s PATH_TO_DATA_FILE", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	defer f.Close()

	if err := day03.Process(f); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}
