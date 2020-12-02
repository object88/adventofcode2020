package day02a

import "testing"

func Test_initProcessor(t *testing.T) {
	p, err := initProcessor()
	if err != nil {
		t.Errorf(err.Error())
	}

	if p.maxIndex == -1 {
		t.Errorf("Failed to find max subexp")
	}
	if p.minIndex == -1 {
		t.Errorf("Failed to find min subexp")
	}
	if p.passwordIndex == -1 {
		t.Errorf("Failed to find password subexp")
	}
	if p.runeIndex == -1 {
		t.Errorf("Failed to find rune subexp")
	}
}

func Test_getSegments(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		min      int
		max      int
		letter   rune
		password string
	}{
		{
			name:     "simple",
			input:    "1-3 a: aa",
			min:      1,
			max:      3,
			letter:   'a',
			password: "aa",
		},
		{
			name:     "simple",
			input:    "1-55 q: abdsjaklfdjskvjkadsf",
			min:      1,
			max:      55,
			letter:   'q',
			password: "abdsjaklfdjskvjkadsf",
		},
	}

	p, _ := initProcessor()

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			min, max, letter, password, err := p.getSegments(tc.input)
			if err != nil {
				t.Errorf(err.Error())
			}
			if min != tc.min {
				t.Errorf("Incorrect min: %d", min)
			}
			if max != tc.max {
				t.Errorf("Incorrect max: %d", max)
			}
			if letter != tc.letter {
				t.Errorf("Incorrect letter: %s", string(letter))
			}
			if password != tc.password {
				t.Errorf("Incorrect password: %s", password)
			}
		})
	}
}
