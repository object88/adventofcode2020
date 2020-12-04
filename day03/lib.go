package day03

import (
	"fmt"
	"io"

	"github.com/object88/adventofcode2020/internal/linereader"
)

type foo struct {
	horzmove int
	vertmove int
	horzpos  int
	vertpos  int
	hits     int
}

func newFoo(horzPath int, vertPath int) foo {
	return foo{horzPath, vertPath, 0, 0, 0}
}

func (f *foo) Hitcount() int {
	return f.hits
}

func (f *foo) Advance(row string) {
	if f.vertpos%f.vertmove == 0 {
		if row[f.horzpos] == '#' {
			f.hits++
		}
		f.horzpos = (f.horzpos + f.horzmove) % len(row)
	}

	f.vertpos++
}

func (f *foo) Report() {
	fmt.Printf("h: %d, v: %d; hits: %d\n", f.horzmove, f.vertmove, f.Hitcount())
}

func Process(r io.Reader) error {
	routes := []foo{
		newFoo(1, 1),
		newFoo(3, 1),
		newFoo(5, 1),
		newFoo(7, 1),
		newFoo(1, 2),
	}

	err := linereader.Read(r, func(raw string) {
		for i := 0; i < len(routes); i++ {
			routes[i].Advance(raw)
		}
	})
	if err != nil {
		return err
	}

	acc := 1
	for i := 0; i < len(routes); i++ {
		routes[i].Report()
		acc *= routes[i].Hitcount()
	}

	fmt.Printf("\nMultiples: %d\n", acc)

	return nil
}
