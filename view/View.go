package view

import (
	"image"
	"strings"

	"github.com/hailongz/golang/canvas"
	"github.com/hailongz/golang/dynamic"
	"golang.org/x/image/font"
)

type ViewUnit struct {
	Scale float64
	Base  float64
}

type IViewProvider interface {
	GetImage(src string) image.Image
	GetFont(v string) font.Face
}

type IView interface {
	CreateElement(name string) IElement
	DeleteElement(id int64)
	GetElement(id int64) IElement
	Set(id int64, key string, value interface{})
	SetSize(width float64, height float64)
	SetUnit(name string, scale float64, base float64)
	Width() float64
	Height() float64
	Calculate(v string, base float64, defaultValue float64) float64
	Layout(id int64)
	Draw(id int64, canvas canvas.Canvas)
	Provider() IViewProvider
}

type View struct {
	autoId     int64
	unitSet    map[string]ViewUnit
	elementSet map[int64]IElement
	width      float64
	height     float64
	provider   IViewProvider
}

func NewView() *View {
	return &View{
		autoId: 0,
		unitSet: map[string]ViewUnit{
			"%":  ViewUnit{Scale: 0, Base: 0.01},
			"px": ViewUnit{Scale: 1, Base: 0},
		},
		elementSet: map[int64]IElement{},
		width:      0,
		height:     0,
	}
}

func (V *View) CreateElement(name string) IElement {
	id := V.autoId + 1
	V.autoId = id
	v := NewElement(id, name)
	V.elementSet[id] = v
	return v
}

func (V *View) DeleteElement(id int64) {
	delete(V.elementSet, id)
}

func (V *View) GetElement(id int64) IElement {
	return V.elementSet[id]
}

func (V *View) Set(id int64, key string, value interface{}) {
	v, ok := V.elementSet[id]
	if ok {
		v.Set(key, value)
	}
}

func (V *View) SetSize(width float64, height float64) {
	V.width = width
	V.height = height
	v, ok := V.elementSet[1]
	if ok {
		e, b := v.(IViewElement)
		if b {
			e.SetWidth(width)
			e.SetHeight(height)
		}
	}
}

func (V *View) SetUnit(name string, scale float64, base float64) {
	V.unitSet[name] = ViewUnit{Scale: scale, Base: base}
}

func (V *View) Width() float64 {
	return V.width
}

func (V *View) Height() float64 {
	return V.height
}

func (V *View) Provider() IViewProvider {
	return V.provider
}

func (V *View) SetProvider(v IViewProvider) {
	V.provider = v
}

func (V *View) Calculate(v string, base float64, defaultValue float64) float64 {

	if v == "" || v == "auto" {
		return defaultValue
	}

	for suffix, unit := range V.unitSet {
		if strings.HasSuffix(v, suffix) {
			r := dynamic.FloatValue(v[0:len(v)-len(suffix)], 0)
			return r*unit.Scale + r*unit.Base*base
		}
	}

	if strings.HasSuffix(v, "vh") {
		r := dynamic.FloatValue(v[0:len(v)-2], 0)
		return r * V.height * 0.01
	}

	if strings.HasSuffix(v, "vw") {
		r := dynamic.FloatValue(v[0:len(v)-2], 0)
		return r * V.width * 0.01
	}

	return dynamic.FloatValue(v, defaultValue)
}

func (V *View) Layout(id int64) {
	v, ok := V.elementSet[id]
	if ok {
		Layout(V, v)
	}
}

func (V *View) Draw(id int64, canvas canvas.Canvas) {
	v, ok := V.elementSet[id]
	if ok {
		Draw(V, canvas, v)
	}
}
