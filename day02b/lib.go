package day02b

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

type processor struct {
	r             *regexp.Regexp
	minIndex      int
	maxIndex      int
	runeIndex     int
	passwordIndex int
}

func initProcessor() (*processor, error) {
	// Expect that the input follows "N-M a: pppp".  Easiest way to process this
	// is probably a regex

	r, err := regexp.Compile(`^(?P<min>[0-9]+)-(?P<max>[0-9]+) (?P<letter>[a-zA-Z]): (?P<password>.+)$`)
	if err != nil {
		return nil, err
	}

	p := &processor{
		r:             r,
		minIndex:      r.SubexpIndex("min"),
		maxIndex:      r.SubexpIndex("max"),
		runeIndex:     r.SubexpIndex("letter"),
		passwordIndex: r.SubexpIndex("password"),
	}
	return p, nil
}

// Process returns the number of valid passwords in the provided input
func Process(in io.Reader) (int, error) {
	p, err := initProcessor()
	if err != nil {
		return -1, err
	}

	// We are going to parse & tokenize all in one go.
	acc := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		raw := scanner.Text()
		if validateLine(p, raw) {
			acc++
		}
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return acc, nil
}

func validateLine(p *processor, in string) bool {
	offset1, offset2, letter, pass, err := p.getSegments(in)
	if err != nil {
		return false
	}

	rune0 := rune(pass[offset1-1])
	rune1 := rune(pass[offset2-1])
	if rune0 == letter && rune1 == letter {
		return false
	} else if rune0 == letter || rune1 == letter {
		return true
	}
	return false
}

func (p *processor) getSegments(in string) (min int, max int, letter rune, pass string, err error) {
	submatches := p.r.FindStringSubmatch(in)
	if len(submatches) == 0 {
		err = fmt.Errorf("Failed to find submatches for '%s'", in)
		return
	}

	if min, err = strconv.Atoi(submatches[p.minIndex]); err != nil {
		return
	}
	if max, err = strconv.Atoi(submatches[p.maxIndex]); err != nil {
		return
	}
	letter = rune(submatches[p.runeIndex][0])
	pass = submatches[p.passwordIndex]

	return
}
