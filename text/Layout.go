package text

type LayoutItem struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
	Loc    int
	Len    int
	Style  *Style
}

type LayoutRow struct {
	Y      float64
	Width  float64
	Height float64
	Loc    int
	Len    int
	Items  []*LayoutItem
}

func (R *LayoutRow) Layout() {
	for _, item := range R.Items {
		if item.Style != nil {
			switch item.Style.VerticalAlignment {
			case VerticalAlignmentMiddle:
				item.Y = (R.Height - item.Height) * 0.5
				break
			case VerticalAlignmentBottom:
				item.Y = (R.Height - item.Height)
				break
			}
		}
	}
}

type Layout struct {
	text     *Text
	maxWidth float64
	width    float64
	height   float64
	Rows     []*LayoutRow
}

func (L *Layout) Text() *Text {
	return L.text
}

func (L *Layout) MaxWidth() float64 {
	return L.maxWidth
}

func (L *Layout) Width() float64 {
	return L.width
}

func (L *Layout) Height() float64 {
	return L.height
}

func NewLayout(text *Text, maxWidth float64, dpi float64, lineSpacing float64) *Layout {

	v := Layout{text: text, maxWidth: maxWidth, Rows: []*LayoutRow{}}

	var x float64 = 0
	var y float64 = 0
	var lineHeight float64 = 0
	var loc int = 0
	var p Index = nil

	row := &LayoutRow{Items: []*LayoutItem{}}

	text.Each(func(vs []rune, idx int, style *Style) int {

		if style.Font == nil {
			return -1
		}

		c := vs[idx]

		if c == '\n' {

			if row.Len > 0 {
				row.Height = lineHeight
				v.Rows = append(v.Rows, row)
				row.Layout()
			}

			if lineHeight == 0 {
				y = y + style.Font.GetDescent() - style.Font.GetAscent()
			} else {
				y = y + lineHeight + lineSpacing
			}
			lineHeight = 0
			x = 0
			p = nil
			loc += 1
			row = &LayoutRow{Items: []*LayoutItem{}}
			row.Loc = loc
			row.Len = 0
			row.Width = 0
			row.Y = y
			return 1
		}

		i := style.Font.Index(vs, idx)

		if i == nil {
			return 1
		}

		n := i.GetLength()

		ih := i.GetFont().GetDescent() - i.GetFont().GetAscent()
		iw := style.Font.Width(i)

		if p != nil {
			iw += style.Font.Kern(p, i)
		}

		p = i

		sep := style.LetterSpacing

		if loc == 0 {
			sep = 0
		}

		if x+iw+sep > maxWidth {

			if row.Len > 0 {
				row.Height = lineHeight
				v.Rows = append(v.Rows, row)
				row.Layout()
			}

			y = y + lineHeight + lineSpacing
			lineHeight = 0
			x = 0
			p = nil
			row = &LayoutRow{Items: []*LayoutItem{}}
			row.Loc = loc
			row.Len = 0
			row.Width = 0
			row.Y = y
		}

		row.Len += n
		row.Width += iw + sep

		if len(row.Items) == 0 {
			row.Items = append(row.Items, &LayoutItem{X: x + sep, Width: iw, Height: ih, Loc: loc, Len: n, Style: style})
		} else {
			l := row.Items[len(row.Items)-1]
			if sep == 0 && l.Style == style {
				l.Len += n
				l.Width += iw + sep
				if l.Height < ih {
					l.Height = ih
				}
			} else {
				row.Items = append(row.Items, &LayoutItem{X: x + sep, Width: iw, Height: ih, Loc: loc, Len: n, Style: style})
			}
		}

		x = x + iw + sep

		if v.width < x {
			v.width = x
		}

		if lineHeight < ih {
			lineHeight = ih
		}

		loc += n

		return n
	})

	v.height = y + lineHeight

	if row.Len > 0 {
		row.Height = lineHeight
		v.Rows = append(v.Rows, row)
		row.Layout()
	}

	return &v
}
