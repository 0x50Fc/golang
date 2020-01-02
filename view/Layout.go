package view

type ILayout interface {
	Layout(view IView)
}

func Layout(view IView, element IElement) {

	v, ok := element.(ILayout)

	if ok {
		v.Layout(view)
	}

}
