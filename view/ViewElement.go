package view

import (
	"math"
	"strings"

	"github.com/hailongz/golang/canvas"
	"github.com/hailongz/golang/dynamic"
)

func init() {
	AddElementConstructor("view", func(id int64, name string) IElement {
		v := &ViewElement{}
		v.SetId(id)
		v.SetName(name)
		return v
	})
}

type IViewElement interface {
	IElement
	X() float64
	Y() float64
	Width() float64
	Height() float64
	ContentWidth() float64
	ContentHeight() float64
	SetX(v float64)
	SetY(v float64)
	SetWidth(v float64)
	SetHeight(v float64)
	SetContentWidth(v float64)
	SetContentHeight(v float64)
}

type ViewElement struct {
	Element
	x             float64
	y             float64
	width         float64
	height        float64
	contentWidth  float64
	contentHeight float64
}

func (E *ViewElement) X() float64 {
	return E.x
}

func (E *ViewElement) Y() float64 {
	return E.y
}

func (E *ViewElement) Width() float64 {
	return E.width
}

func (E *ViewElement) Height() float64 {
	return E.height
}

func (E *ViewElement) ContentWidth() float64 {
	return E.contentWidth
}

func (E *ViewElement) ContentHeight() float64 {
	return E.contentHeight
}

func (E *ViewElement) SetX(v float64) {
	E.x = v
}

func (E *ViewElement) SetY(v float64) {
	E.y = v
}

func (E *ViewElement) SetWidth(v float64) {
	E.width = v
}

func (E *ViewElement) SetHeight(v float64) {
	E.height = v
}

func (E *ViewElement) SetContentWidth(v float64) {
	E.contentWidth = v
}

func (E *ViewElement) SetContentHeight(v float64) {
	E.contentHeight = v
}

func relative(E *ViewElement, view IView) {

	p_top, p_right, p_bottom, p_left := Edge(E, view, "padding")

	inWidth := E.Width() - p_left - p_right

	inHeight := E.Height() - p_top - p_bottom

	p := E.FirstChild()

	E.SetContentWidth(0)
	E.SetContentHeight(0)

	for p != nil {

		e, ok := p.(IViewElement)

		if ok {

			e.SetWidth(view.Calculate(dynamic.StringValue(p.Get("width"), "auto"), inWidth, math.MaxFloat64))
			e.SetHeight(view.Calculate(dynamic.StringValue(p.Get("height"), "auto"), inHeight, math.MaxFloat64))

			Layout(view, p)

			m_top, m_right, m_bottom, m_left := Edge(e, view, "margin")

			left := view.Calculate(dynamic.StringValue(p.Get("left"), "auto"), inWidth, math.MaxFloat64)
			right := view.Calculate(dynamic.StringValue(p.Get("right"), "auto"), inWidth, math.MaxFloat64)
			top := view.Calculate(dynamic.StringValue(p.Get("top"), "auto"), inHeight, math.MaxFloat64)
			bottom := view.Calculate(dynamic.StringValue(p.Get("bottom"), "auto"), inHeight, math.MaxFloat64)

			if left == math.MaxFloat64 {
				if right == math.MaxFloat64 {
					left = p_left + m_left
				} else {
					left = p_left + inWidth - m_right - right - e.Width()
				}
			} else {
				left = p_left + left + m_left
			}

			if top == math.MaxFloat64 {
				if bottom == math.MaxFloat64 {
					top = p_top + m_top
				} else {
					top = p_top + inHeight - m_bottom - bottom - e.Height()
				}
			} else {
				top = p_top + top + m_top
			}

			e.SetX(left)
			e.SetY(top)

			if left+e.Width()+m_right+p_right > E.ContentWidth() {
				E.SetContentWidth(left + e.Width() + m_right + p_right)
			}

			if top+e.Height()+m_bottom+p_bottom > E.ContentHeight() {
				E.SetContentHeight(top + e.Height() + m_bottom + p_bottom)
			}
		}

		p = p.NextSibling()

	}

	if E.Width() == math.MaxFloat64 {
		E.SetWidth(E.ContentWidth())
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
		E.SetHeight(E.ContentHeight())
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

type flexElement struct {
	element IViewElement
	left    float64
	top     float64
	right   float64
	bottom  float64
}

func flex(E *ViewElement, view IView) {

	p_top, p_right, p_bottom, p_left := Edge(E, view, "padding")

	inWidth := E.Width() - p_left - p_right

	inHeight := E.Height() - p_top - p_bottom

	p := E.FirstChild()

	x := p_left
	y := p_top

	lineHeight := float64(0)
	maxWidth := p_left
	lineElements := []*flexElement{}

	alignItems := dynamic.StringValue(E.Get("align-items"), "flex-start")
	justifyContent := dynamic.StringValue(E.Get("justify-content"), "flex-start")

	row := func(top float64, lineHeight float64) {
		n := len(lineElements)
		if n > 0 {

			{
				e := lineElements[n-1]
				useWidth := e.element.X() + e.element.Width() + e.right - p_left

				if useWidth < inWidth {

					dv := inWidth - useWidth

					grow := int64(0)

					for _, e = range lineElements {
						grow = grow + dynamic.IntValue(e.element.Get("flex-grow"), 0)
					}

					if grow > 0 {

						for _, e = range lineElements {
							v := dynamic.IntValue(e.element.Get("flex-grow"), 0)
							if v > 0 {
								vv := (float64(v) * dv / float64(grow))
								useWidth = useWidth + vv
								e.element.SetWidth(e.element.Width() + vv)
								Layout(view, e.element)
							}
						}

					}
				}

				if useWidth < inWidth {
					switch justifyContent {
					case "flex-end":
						{
							dv := inWidth - useWidth
							for _, e = range lineElements {
								e.element.SetX(e.element.X() + dv)
							}
						}
						break
					case "center":
						{
							dv := (inWidth - useWidth) * 0.5
							for _, e = range lineElements {
								e.element.SetX(e.element.X() + dv)
							}
						}
						break
					}
				}

			}

			for _, e := range lineElements {
				switch alignItems {
				case "flex-end":
					{
						e.element.SetY(top + lineHeight - e.element.Height() - e.bottom)
					}
					break
				case "center":
					{
						e.element.SetY(top + e.top + (lineHeight-e.top-e.bottom-e.element.Height())*0.5)
					}
					break
				case "stretch":
					{
						e.element.SetY(top + e.top)
						e.element.SetHeight((lineHeight - e.top - e.bottom))
						Layout(view, e.element)
					}
					break
				}
			}

			lineElements = []*flexElement{}
		}
	}

	for p != nil {

		e, ok := p.(IViewElement)

		if ok && dynamic.StringValue(e.Get("display"), "") != "none" {

			e.SetWidth(view.Calculate(dynamic.StringValue(e.Get("width"), "auto"), inWidth, math.MaxFloat64))
			e.SetHeight(view.Calculate(dynamic.StringValue(e.Get("height"), "auto"), inHeight, math.MaxFloat64))

			Layout(view, p)

			m_top, m_right, m_bottom, m_left := Edge(e, view, "margin")

			if x+m_left+e.Width()+m_right > p_left+inWidth {
				if len(lineElements) > 0 {
					row(y, lineHeight)
				}
				y = y + lineHeight
				x = p_left
				lineHeight = 0
			}

			e.SetX(x + m_left)
			e.SetY(y + m_top)

			x = x + m_left + e.Width() + m_right

			if x > maxWidth {
				maxWidth = x
			}

			if m_top+e.Height()+m_bottom > lineHeight {
				lineHeight = m_top + e.Height() + m_bottom
			}

			lineElements = append(lineElements, &flexElement{element: e, left: m_left, right: m_right, top: m_top, bottom: m_bottom})

		}

		p = p.NextSibling()
	}

	if len(lineElements) > 0 {
		row(y, lineHeight)
	}

	y = y + lineHeight

	E.SetContentWidth(maxWidth + p_right)
	E.SetContentHeight(y + p_bottom)

	if E.Width() == math.MaxFloat64 {
		E.SetWidth(E.ContentWidth())
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
		E.SetHeight(E.ContentHeight())
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

func (E *ViewElement) Layout(view IView) {

	switch dynamic.StringValue(E.Get("display"), "relative") {
	case "flex":
		flex(E, view)
		break
	default:
		relative(E, view)
		break
	}
}

func Edge(e IElement, view IView, key string) (float64, float64, float64, float64) {

	var top, right, bottom, left, width, height float64 = 0, 0, 0, 0, 0, 0

	{
		E, ok := e.(IViewElement)
		if ok {
			width = E.Width()
			height = E.Height()
		}
	}

	v := dynamic.StringValue(e.Get(key), "")

	if v != "" {
		vs := strings.Split(v, " ")
		n := len(vs)
		if n > 0 {
			top = view.Calculate(vs[0], height, 0)
			if n > 1 {
				right = view.Calculate(vs[1], width, 0)
				if n > 2 {
					bottom = view.Calculate(vs[2], height, 0)
					if n > 3 {
						left = view.Calculate(vs[3], width, 0)
					} else {
						left = right
					}
				} else {
					bottom = top
					left = right
				}
			} else {
				right = top
				bottom = top
				left = top
			}
		}
	}

	return top, right, bottom, left
}

func (E *ViewElement) Draw(view IView, canvas canvas.Canvas) {

	r := view.Calculate(dynamic.StringValue(E.Get("border-radius"), ""), 0, 0)

	{
		v := E.Get("background-color")
		if v != nil {
			canvas.Save()
			canvas.SetHexColor(dynamic.StringValue(v, ""))
			if r > 0 {
				canvas.RoundedRectangle(0, 0, E.Width(), E.Height(), r)
			} else {
				canvas.Rectangle(0, 0, E.Width(), E.Height())
			}
			canvas.Fill()
			canvas.Restore()
		}
	}

	{
		v := E.Get("background-image")
		if v != nil {
			provider := view.Provider()
			if provider != nil {
				im := provider.GetImage(dynamic.StringValue(v, ""))
				if im != nil {

					size := im.Bounds().Size()

					if size.X > 0 && size.Y > 0 && E.Width() > 0 && E.Height() > 0 {

						canvas.Save()

						if r > 0 {
							canvas.RoundedRectangle(0, 0, E.Width(), E.Height(), r)
						} else {
							canvas.Rectangle(0, 0, E.Width(), E.Height())
						}

						canvas.Clip()

						ws := E.Width() / float64(size.X)
						hs := E.Height() / float64(size.Y)

						if ws < hs {
							canvas.Scale(hs, hs)
							sw := float64(size.X) * hs
							canvas.DrawImage(im, int((E.Width()-sw)*0.5/hs), 0)
						} else {
							canvas.Scale(ws, ws)
							sh := float64(size.Y) * ws
							canvas.DrawImage(im, 0, int((E.Height()-sh)*0.5/ws))
						}

						canvas.Restore()
					}

				}
			}

		}
	}

	{
		width := view.Calculate(dynamic.StringValue(E.Get("border-width"), ""), 0, 0)
		v := E.Get("border-color")
		if v != nil && width > 0 {

			canvas.Save()
			canvas.SetHexColor(dynamic.StringValue(v, ""))
			canvas.SetLineWidth(width)
			if r > 0 {
				canvas.RoundedRectangle(0, 0, E.Width(), E.Height(), r)
			} else {
				canvas.Rectangle(0, 0, E.Width(), E.Height())
			}
			canvas.Stroke()
			canvas.Restore()
		}
	}
}
