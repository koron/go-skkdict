package skkdict

import (
	"bytes"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	f := func(s string, expected []*Entry) {
		br := bytes.NewReader([]byte(s))
		r := NewReader(br)
		for _, exp := range expected {
			act, err := r.Read()
			if err != nil {
				t.Fatalf("Reader.Read failed: %s", err)
			}
			assertEquals(t, act, exp, "source:%q", r.ll)
		}
		_, err := r.Read()
		if err == nil {
			t.Fatalf("Reader should end with io.EOF but got line:%q", r.ll)
		} else if err != io.EOF {
			t.Fatalf("Reader failed unexpectedly: %s", err)
		}
	}

	f(``, []*Entry{})
	f(`;;`, []*Entry{nil})
	f(`りりs /凛々/凛凛/律々;=凛々しい/`, []*Entry{
		{
			Label: "りりs",
			Words: []Word{
				{Text: "凛々"},
				{Text: "凛凛"},
				{Text: "律々", Desc: "=凛々しい"},
			},
		},
	})
	f(`;;
あい /愛/哀/相/挨/
おん /音/温/御/恩/穏/遠/
;;
ぐん /群/郡/軍/
こと /事/異/言/殊/琴/`, []*Entry{
		nil,
		{
			Label: "あい",
			Words: []Word{
				{Text: "愛"},
				{Text: "哀"},
				{Text: "相"},
				{Text: "挨"},
			},
		},
		{
			Label: "おん",
			Words: []Word{
				{Text: "音"},
				{Text: "温"},
				{Text: "御"},
				{Text: "恩"},
				{Text: "穏"},
				{Text: "遠"},
			},
		},
		nil,
		{
			Label: "ぐん",
			Words: []Word{
				{Text: "群"},
				{Text: "郡"},
				{Text: "軍"},
			},
		},
		{
			Label: "こと",
			Words: []Word{
				{Text: "事"},
				{Text: "異"},
				{Text: "言"},
				{Text: "殊"},
				{Text: "琴"},
			},
		},
	})

}
