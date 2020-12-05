package wordreader

import (
	"bufio"
	"io"
	"strings"
)

func Read(r io.Reader, f0 func(string), f1 func()) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := scanner.Text()
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			f1()
		} else {
			for _, seg := range strings.Split(l, " ") {
				f0(seg)
			}
		}
	}

	f1()

	// if err := scanner.Err(); err != nil {
	// 	return err
	// }

	// return nil

	// reader := bufio.NewReaderSize(r, 1024)
	// var buf strings.Builder
	// for {
	// 	line := reader.ReadLine()

	// 	rn, _, err := reader.ReadRune()
	// 	if unicode.IsSpace(rn) {
	// 		if buf.Len() == 0 {
	// 			f1()
	// 			buf.Reset()
	// 		} else {
	// 			f0(buf.String())
	// 		}
	// 	} else {
	// 		buf.WriteRune(rn)
	// 	}

	// 	if err == io.EOF {
	// 		f1()
	// 		return
	// 	}
	// }
}
