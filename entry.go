package skkdict

// Word represents a word.
type Word struct {
	Text string
	Desc string
}

// Entry represents an entry in SKK dictionary.
type Entry struct {
	Label string
	Words []Word
}
