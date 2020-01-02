package emoji

import (
	"image"
	"os"
)

type Index interface {
	GetImage() (image.Image, error)
	GetRune() []rune
	String() string
}

type node struct {
	v   rune
	vs  []rune
	p   string
	src image.Image
	m   map[rune]*node
}

func newNode(v rune) *node {
	return &node{v: v, m: map[rune]*node{}}
}

func (n *node) GetImage() (image.Image, error) {

	if n.src == nil {
		fd, err := os.Open(n.p)
		if err != nil {
			return nil, err
		}
		defer fd.Close()
		n.src, _, err = image.Decode(fd)
		if err != nil {
			return nil, err
		}
	}

	return n.src, nil
}

func (n *node) GetRune() []rune {
	return n.vs
}

func (n *node) String() string {
	return string(n.vs)
}

func (n *node) add(vs []rune, p string, idx int) {

	v := vs[idx]
	s, ok := n.m[v]

	if !ok {
		s = newNode(v)
		n.m[v] = s
	}

	idx = idx + 1

	if idx < len(vs) {
		s.add(vs, p, idx)
	} else {
		s.p = p
		s.vs = vs
	}

}

func (n *node) index(vs []rune, idx int) Index {

	v := vs[idx]
	s, ok := n.m[v]

	if ok {

		idx = idx + 1

		if idx < len(vs) {
			i := s.index(vs, idx)
			if i != nil {
				return i
			}
		}

		if s.vs != nil {
			return s
		}

	}

	return nil
}

type Emoji struct {
	root node
}

func NewEmoji() *Emoji {
	v := &Emoji{}
	v.root.v = -1
	v.root.m = map[rune]*node{}
	return v
}

func (E *Emoji) Add(vs []rune, p string) {
	E.root.add(vs, p, 0)
}

func (E *Emoji) Index(vs []rune, idx int) Index {
	return E.root.index(vs, idx)
}
