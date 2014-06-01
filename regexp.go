package filetask

import (
	"bufio"
	"bytes"
	"regexp"
)

type RegexpTask struct {
	Regexp  *regexp.Regexp
	Convert func([][]byte) [][]byte
	Writer  io.Writer
	Reader  io.Reader
}

func (rt *RegexpTask) Do() error {
	scanner := bufio.NewScanner(rt.Reader)
	for scanner.Scan() {
		newLine := rt.Regexp.ReplaceAllFunc(scanner.Bytes(), func(match []byte) []byte {
			matches := rt.Regexp.FindSubmatch(match)
			return bytes.Join(rt.Convert(matches[1:]), nil)
		})
		if _, err := rt.Writer.Write(append(newLine, '\n')); err != nil {
			return err
		}
	}
	return scanner.Err()
}
