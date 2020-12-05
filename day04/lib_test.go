package day04

import (
	"strings"
	"testing"

	"github.com/object88/adventofcode2020/internal/wordreader"
)

func Test0(t *testing.T) {
	raw := `iyr:2010 ecl:gry hgt:181cm
	pid:591597745 byr:1920 hcl:#6b5442 eyr:2029 cid:123`
	r := strings.NewReader(raw)
	p := &passport{}
	wordreader.Read(r, p.parse, p.foo)
}
