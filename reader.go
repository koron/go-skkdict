package skkdict

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type lineReader interface {
	ReadLine() (line []byte, isPrefix bool, err error)
}

// Reader reads SKK dictionary file and parses as Entry.
type Reader struct {
	lr   lineReader
	lnum int64
	ll   string
}

// NewReader creates a SKK dictionary Reader.
func NewReader(r io.Reader) *Reader {
	lr, ok := r.(lineReader)
	if !ok {
		lr = bufio.NewReader(r)
	}
	return &Reader{
		lr: lr,
	}
}

// Read reads an entry.
func (r *Reader) Read() (*Entry, error) {
	r.ll = ""
	b, isPrefix, err := r.lr.ReadLine()
	if err != nil {
		return nil, err
	}
	r.lnum++
	if isPrefix {
		return nil, fmt.Errorf("too long line at %d", r.lnum)
	}
	r.ll = string(b)
	return parseEntry(r.ll)
}

// parseEntry parses a string as Entry.
func parseEntry(s string) (*Entry, error) {
	if strings.HasPrefix(s, ";;") {
		return nil, nil
	}
	s = strings.TrimRight(s, " \t\r\n")
	items := strings.SplitN(s, " ", 2)
	if items == nil || len(items) != 2 {
		return nil, fmt.Errorf("Invalid format: %s", s)
	}
	label := items[0]
	values := strings.Split(strings.Trim(items[1], "/"), "/")
	words := make([]Word, len(values))
	for i, v := range values {
		words[i] = parseWord(v)
	}
	return &Entry{
		Label: label,
		Words: words,
	}, nil
}

// parseWord parses a string as Word
func parseWord(v string) Word {
	n := strings.Index(v, ";")
	if n < 0 {
		return Word{Text: v}
	}
	return Word{
		Text: v[0:n],
		Desc: v[n+1:],
	}
}
