package text

import (
	"bytes"
)

type Item struct {
	Style *Style
	Text  string
}

type Text struct {
	items []*Item
}

func NewText() *Text {
	return &Text{items: []*Item{}}
}

func (T *Text) Add(text string, style *Style) {
	T.items = append(T.items, &Item{Style: style, Text: text})
}

func (T *Text) Each(fn func(vs []rune, i int, style *Style) int) {
	for _, item := range T.items {
		i := 0
		vs := []rune(item.Text)
		n := len(vs)
		for i < n {
			e := fn(vs, i, item.Style)
			if e > 0 {
				i = i + e
			} else {
				break
			}
		}
	}
}

func (T *Text) StringWithRange(loc int, length int) string {
	b := bytes.NewBuffer(nil)

	i := 0

	for _, item := range T.items {
		bs := []rune(item.Text)
		n := len(bs)
		if i+n > loc {
			if i+n < loc+length {
				c := n - loc - i
				b.WriteString(string(bs[loc-i : loc-i+c]))
				loc += c
				length -= c
			} else {
				b.WriteString(string(bs[loc-i : loc-i+length]))
				loc += length
				length = 0
			}
			if length <= 0 {
				break
			}
		}
		i += n
	}

	return b.String()
}
