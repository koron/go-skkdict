# SKK dictionary reader

[![GoDoc](https://godoc.org/github.com/koron/go-skkdict?status.svg)](https://godoc.org/github.com/koron/go-skkdict)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron/go-skkdict)](https://goreportcard.com/report/github.com/koron/go-skkdict)
[![Build Status](https://travis-ci.org/koron/go-skkdict.svg?branch=master)](https://travis-ci.org/koron/go-skkdict)

Example

```go
import skkdict "github.com/koron/go-skkdict"

f, err := os.Open("SKK-JISYO.utf-8.S")
if err != nil {
	panic(err)
}
defer f.Close()

r := skkdict.NewReader(f)
for {
	entry, err := r.Read()
	if err != nil {
		break
	}
	fmt.Printf("%+v\n", entry)
}
```
