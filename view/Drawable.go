package view

import (
	"github.com/hailongz/golang/canvas"
	"github.com/hailongz/golang/dynamic"
)

type Drawable interface {
	Draw(view IView, canvas canvas.Canvas)
}

func Draw(view IView, canvas canvas.Canvas, element IElement) {

	if dynamic.StringValue(element.Get("display"), "") == "none" {
		return
	}

	{
		v, ok := element.(IViewElement)

		if ok {

			canvas.Translate(v.X(), v.Y())

			if dynamic.StringValue(element.Get("overflow"), "") == "hidden" {
				r := view.Calculate(dynamic.StringValue(element.Get("border-radius"), ""), 0, 0)
				if r > 0 {
					canvas.RoundedRectangle(0, 0, v.Width(), v.Height(), r)
				} else {
					canvas.Rectangle(0, 0, v.Width(), v.Height())
				}
				canvas.Clip()
			}

			{
				draw, ok := element.(Drawable)
				if ok {
					canvas.Save()
					draw.Draw(view, canvas)
					canvas.Restore()
				}
			}

			p := element.FirstChild()

			for p != nil {
				canvas.Save()
				Draw(view, canvas, p)
				canvas.Restore()
				p = p.NextSibling()
			}

			canvas.ResetClip()
		}
	}

}
