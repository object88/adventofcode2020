package main

import (
	"fmt"
	"os"

	"github.com/object88/adventofcode2020/day01b"
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

	x, err := day01b.Process(f, 2020)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%d\n", x)
}
