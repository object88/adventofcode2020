package linereader

import (
	"bufio"
	"io"
)

// Read will read from the provider io.Reader, and feed the contents to the
// channel, one line at a time.  Read will return an error only if there is a
// problem consuming the io.Reader.  Read will close the channel when there is
// nothing left to process.
func Read(r io.Reader, f func(string)) error {
	// We are going to parse & tokenize all in one go.
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
