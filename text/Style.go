package text

import (
	"strconv"
	"strings"
)

const (
	TextAlignmentLeft      = 0
	TextAlignmentCenter    = 1
	TextAlignmentRight     = 2
	TextAlignmentJustified = 3
)

type TextAlignment int

const (
	VerticalAlignmentTop    = 0
	VerticalAlignmentMiddle = 1
	VerticalAlignmentBottom = 2
)

type VerticalAlignment int

type Style struct {
	Font              Font
	Color             uint32
	VerticalAlignment VerticalAlignment
	LetterSpacing     float64
}

func (T TextAlignment) String() string {
	switch T {
	case TextAlignmentCenter:
		return "center"
	case TextAlignmentRight:
		return "right"
	case TextAlignmentJustified:
		return "justified"
	}
	return "left"
}

func TextAlignmentFromString(v string) TextAlignment {
	switch v {
	case "center":
		return TextAlignmentCenter
	case "right":
		return TextAlignmentRight
	case "justified":
		return TextAlignmentJustified
	}
	return TextAlignmentLeft
}

func VerticalAlignmentFromString(v string) VerticalAlignment {
	switch v {
	case "middle":
		return VerticalAlignmentMiddle
	case "bottom":
		return VerticalAlignmentBottom
	}
	return VerticalAlignmentTop
}

func ColorFromString(v string) uint32 {

	if strings.HasPrefix(v, "#") {
		var r, g, b, a uint64 = 0, 0, 0, 0xff
		c := len(v)
		if c == 9 {
			a, _ = strconv.ParseUint(v[1:3], 16, 32)
			r, _ = strconv.ParseUint(v[3:5], 16, 32)
			g, _ = strconv.ParseUint(v[5:7], 16, 32)
			b, _ = strconv.ParseUint(v[7:9], 16, 32)
		} else if c == 7 {
			r, _ = strconv.ParseUint(v[1:3], 16, 32)
			g, _ = strconv.ParseUint(v[3:5], 16, 32)
			b, _ = strconv.ParseUint(v[5:7], 16, 32)
		} else if c == 4 {
			r, _ = strconv.ParseUint(v[1:2], 16, 32)
			g, _ = strconv.ParseUint(v[2:3], 16, 32)
			b, _ = strconv.ParseUint(v[3:4], 16, 32)
			r = (r << 4) | r
			g = (g << 4) | g
			b = (b << 4) | b
		}
		return uint32((a << 24) | (r << 16) | (g << 8) | b)
	}

	return 0
}
