package day04

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/object88/adventofcode2020/internal/wordreader"
)

type passport struct {
	count int
	ecl   string
	eyr   string
	hcl   string
	byr   string
	iyr   string
	pid   string
	cid   string
	hgt   string
}

func (p *passport) parse(in string) {
	segments := strings.SplitN(in, ":", 2)
	switch segments[0] {
	case "ecl":
		p.ecl = segments[1]
	case "eyr":
		p.eyr = segments[1]
	case "hcl":
		p.hcl = segments[1]
	case "byr":
		p.byr = segments[1]
	case "iyr":
		p.iyr = segments[1]
	case "pid":
		p.pid = segments[1]
	case "cid":
		p.cid = segments[1]
	case "hgt":
		p.hgt = segments[1]
	}
}

func (p *passport) foo() {
	if p.isValid() {
		p.count++
	}
	p.reset()
}

func (p *passport) isValid() bool {
	if !yearrange(p.byr, 1920, 2002) {
		return false
	}
	if !yearrange(p.iyr, 2010, 2020) {
		return false
	}
	if !yearrange(p.eyr, 2020, 2030) {
		return false
	}

	if strings.HasSuffix(p.hgt, "in") {
		if n, err := strconv.Atoi(p.hgt[:len(p.hgt)-2]); err != nil {
			return false
		} else if !inrange(n, 59, 76) {
			return false
		}
	} else if strings.HasSuffix(p.hgt, "cm") {
		if n, err := strconv.Atoi(p.hgt[:len(p.hgt)-2]); err != nil {
			return false
		} else if !inrange(n, 150, 193) {
			return false
		}
	} else {
		return false
	}

	if len(p.hcl) == 0 {
		return false
	} else if p.hcl[0] != '#' {
		return false
	} else if len(p.hcl) != 7 {
		return false
	}
	for i := 1; i < 7; i++ {
		if !inrange(int(p.hcl[i]), '0', '9') && !inrange(int(p.hcl[i]), 'a', 'f') {
			return false
		}
	}

	switch p.ecl {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
	default:
		return false
	}

	if len(p.pid) != 9 {
		return false
	}
	for i := 0; i < 9; i++ {
		if !inrange(int(p.pid[i]), '0', '9') {
			return false
		}
	}

	return true
}

func yearrange(in string, min, max int) bool {
	val, err := strconv.Atoi(in)
	if err != nil {
		return false
	}
	return inrange(val, min, max)
}

func inrange(val, min, max int) bool {
	return val >= min && val <= max
}

func (p *passport) reset() {
	p.ecl = ""
	p.eyr = ""
	p.hcl = ""
	p.byr = ""
	p.iyr = ""
	p.pid = ""
	p.cid = ""
	p.hgt = ""
}

func Process(r io.Reader) error {
	p := &passport{}

	wordreader.Read(r, p.parse, p.foo)

	fmt.Printf("valid count: %d\n", p.count)

	return nil
}
