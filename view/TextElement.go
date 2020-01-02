package view

import (
	"image"
	"math"

	"github.com/hailongz/golang/canvas"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/text"
)

func init() {
	AddElementConstructor("text", func(id int64, name string) IElement {
		v := &TextElement{}
		v.SetId(id)
		v.SetName(name)
		return v
	})

}

type TextElement struct {
	ViewElement
	text *text.Layout
}

func (E *TextElement) Layout(view IView) {

	p_top, p_right, p_bottom, p_left := Edge(E, view, "padding")

	E.text = nil
	E.SetContentWidth(0)
	E.SetContentHeight(0)

	maxWidth := E.Width()

	if maxWidth == math.MaxFloat64 {
		maxWidth = view.Calculate(dynamic.StringValue(E.Get("max-width"), "auto"), 0, math.MaxFloat64)
	}

	if maxWidth != math.MaxFloat64 {
		maxWidth = maxWidth - p_left - p_right
	}

	{
		t := text.NewText()
		v := E.Get("#text")

		def := text.Style{}
		def.Font, _ = text.NewFont(dynamic.StringValue(E.Get("font-family"), ""), view.Calculate(dynamic.StringValue(E.Get("font-size"), "14"), 0, 0))
		def.LetterSpacing = view.Calculate(dynamic.StringValue(E.Get("letter-spacing"), ""), 0, 0)
		def.Color = text.ColorFromString(dynamic.StringValue(E.Get("color"), "#000000"))
		def.VerticalAlignment = text.VerticalAlignmentFromString(dynamic.StringValue(E.Get("vertical-align"), ""))

		if v != nil {

			t.Add(dynamic.StringValue(v, ""), &def)
		}

		{
			p := E.FirstChild()

			for p != nil {

				v = p.Get("#text")

				if v != nil {

					s := text.Style{}
					s.Font, _ = text.NewFont(dynamic.StringValue(p.Get("font-family"), dynamic.StringValue(E.Get("font-family"), "")), view.Calculate(dynamic.StringValue(p.Get("font-size"), dynamic.StringValue(E.Get("font-size"), "14")), 0, 0))
					s.Color = text.ColorFromString(dynamic.StringValue(p.Get("color"), dynamic.StringValue(E.Get("color"), "#000000")))
					s.LetterSpacing = view.Calculate(dynamic.StringValue(p.Get("letter-spacing"), dynamic.StringValue(E.Get("letter-spacing"), "")), 0, 0)
					s.VerticalAlignment = text.VerticalAlignmentFromString(dynamic.StringValue(p.Get("vertical-align"), dynamic.StringValue(E.Get("vertical-align"), "")))

					t.Add(dynamic.StringValue(v, ""), &s)

				}

				p = p.NextSibling()

			}
		}

		E.text = text.NewLayout(t, maxWidth, 72, view.Calculate(dynamic.StringValue(E.Get("line-spacing"), "auto"), 0, 0))
		E.SetContentWidth(E.text.Width())
		E.SetContentHeight(E.text.Height())
	}

	if E.Width() == math.MaxFloat64 {
		E.SetWidth(E.ContentWidth() + p_left + p_right)
		min := view.Calculate(dynamic.StringValue(E.Get("min-width"), "auto"), 0, 0)
		max := view.Calculate(dynamic.StringValue(E.Get("max-width"), "auto"), 0, math.MaxFloat64)
		if E.Width() < min {
			E.SetWidth(min)
		}
		if E.Width() > max {
			E.SetWidth(max)
		}
	}

	if E.Height() == math.MaxFloat64 {
		E.SetHeight(E.ContentHeight() + p_top + p_bottom)
		min := view.Calculate(dynamic.StringValue(E.Get("min-height"), "auto"), 0, 0)
		max := view.Calculate(dynamic.StringValue(E.Get("max-height"), "auto"), 0, math.MaxFloat64)
		if E.Height() < min {
			E.SetHeight(min)
		}
		if E.Height() > max {
			E.SetHeight(max)
		}
	}
}

func (E *TextElement) Draw(view IView, canvas canvas.Canvas) {
	E.ViewElement.Draw(view, canvas)

	if E.text != nil {

		p_top, p_right, _, p_left := Edge(E, view, "padding")

		w := int(math.Ceil(E.Width()))
		h := int(math.Ceil(E.Height()))

		if w > 0 && h > 0 {

			v := image.NewRGBA(image.Rect(0, 0, w, h))

			text.Draw(v, E.text, 0, p_top, E.Width(), E.Height(), p_left, p_right, text.TextAlignmentFromString(dynamic.StringValue(E.Get("text-align"), "left")))

			canvas.DrawImage(v, 0, 0)

		}

	}
}
