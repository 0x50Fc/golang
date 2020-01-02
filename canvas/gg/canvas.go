package gg

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
	"gopkg.in/fogleman/gg.v1"
)

type Canvas struct {
	ctx *gg.Context
}

func NewCanvas(width, height int) *Canvas {
	return &Canvas{ctx: gg.NewContext(width, height)}
}

func (C *Canvas) Image() image.Image {
	return C.ctx.Image()
}

func (C *Canvas) Width() int {
	return C.ctx.Width()
}

func (C *Canvas) Height() int {
	return C.ctx.Height()
}

func (C *Canvas) Clear() {
	C.ctx.Clear()
}

func (C *Canvas) ClearPath() {
	C.ctx.ClearPath()
}

func (C *Canvas) Clip() {
	C.ctx.Clip()
}

func (C *Canvas) ClosePath() {
	C.ctx.ClosePath()
}

func (C *Canvas) BeginPath() {
	C.ctx.NewSubPath()
}

func (C *Canvas) CubicTo(x1, y1, x2, y2, x3, y3 float64) {
	C.ctx.CubicTo(x1, y1, x2, y2, x3, y3)
}

func (C *Canvas) Arc(x, y, r, angle1, angle2 float64) {
	C.ctx.DrawArc(x, y, r, angle1, angle2)
}

func (C *Canvas) Circle(x, y, r float64) {
	C.ctx.DrawCircle(x, y, r)
}

func (C *Canvas) Ellipse(x, y, rx, ry float64) {
	C.ctx.DrawEllipse(x, y, rx, ry)
}

func (C *Canvas) Line(x1, y1, x2, y2 float64) {
	C.ctx.DrawLine(x1, y1, x2, y2)
}

func (C *Canvas) Rectangle(x, y, w, h float64) {
	C.ctx.DrawRectangle(x, y, w, h)
}

func (C *Canvas) RoundedRectangle(x, y, w, h, r float64) {
	C.ctx.DrawRoundedRectangle(x, y, w, h, r)
}

func (C *Canvas) DrawImage(im image.Image, x, y int) {
	C.ctx.DrawImage(im, x, y)
}

func (C *Canvas) DrawString(s string, x, y float64) {
	C.ctx.DrawString(s, x, y)
}

func (C *Canvas) Fill() {
	C.ctx.Fill()
}

func (C *Canvas) LineTo(x, y float64) {
	C.ctx.LineTo(x, y)
}

func (C *Canvas) SetFont(v font.Face) {
	C.ctx.SetFontFace(v)
}

func (C *Canvas) MeasureString(s string) (w, h float64) {
	return C.ctx.MeasureString(s)
}

func (C *Canvas) MoveTo(x, y float64) {
	C.ctx.MoveTo(x, y)
}

func (C *Canvas) Save() {
	C.ctx.Push()
}

func (C *Canvas) Restore() {
	C.ctx.Pop()
}

func (C *Canvas) ResetClip() {
	C.ctx.ResetClip()
}

func (C *Canvas) QuadraticTo(x1, y1, x2, y2 float64) {
	C.ctx.QuadraticTo(x1, y1, x2, y2)
}

func (C *Canvas) Scale(x, y float64) {
	C.ctx.Scale(x, y)
}

func (C *Canvas) SetColor(c color.Color) {
	C.ctx.SetColor(c)
}

func (C *Canvas) SetDash(dashes ...float64) {
	C.ctx.SetDash(dashes...)
}

func (C *Canvas) SetDashOffset(offset float64) {
	C.ctx.SetDashOffset(offset)
}

func (C *Canvas) SetHexColor(x string) {
	C.ctx.SetHexColor(x)
}

func (C *Canvas) SetLineWidth(lineWidth float64) {
	C.ctx.SetLineWidth(lineWidth)
}

func (C *Canvas) Stroke() {
	C.ctx.Stroke()
}

func (C *Canvas) Rotate(angle float64) {
	C.ctx.Rotate(angle)
}

func (C *Canvas) Translate(x, y float64) {
	C.ctx.Translate(x, y)
}
