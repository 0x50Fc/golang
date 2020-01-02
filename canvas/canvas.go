package canvas

import (
	"bytes"
	"image"
	"image/color"
	"io/ioutil"
	"sync"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type Canvas interface {
	Image() image.Image
	Width() int
	Height() int
	Clear()
	ClearPath()
	Clip()
	ResetClip()
	ClosePath()
	BeginPath()
	CubicTo(x1, y1, x2, y2, x3, y3 float64)
	Arc(x, y, r, angle1, angle2 float64)
	Circle(x, y, r float64)
	Ellipse(x, y, rx, ry float64)
	Line(x1, y1, x2, y2 float64)
	Rectangle(x, y, w, h float64)
	RoundedRectangle(x, y, w, h, r float64)
	DrawImage(im image.Image, x, y int)
	DrawString(s string, x, y float64)
	Fill()
	LineTo(x, y float64)
	SetFont(v font.Face)
	MeasureString(s string) (w, h float64)
	MoveTo(x, y float64)
	Save()
	Restore()
	QuadraticTo(x1, y1, x2, y2 float64)
	Scale(x, y float64)
	SetColor(c color.Color)
	SetDash(dashes ...float64)
	SetDashOffset(offset float64)
	SetHexColor(x string)
	SetLineWidth(lineWidth float64)
	Stroke()
	Rotate(angle float64)
	Translate(x, y float64)
}

var lock sync.RWMutex
var fonts = map[string]*truetype.Font{}

func LoadFont(path string) (*truetype.Font, error) {

	lock.RLock()

	v, ok := fonts[path]

	lock.RUnlock()

	if ok {
		return v, nil
	}

	b, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	v, err = freetype.ParseFont(b)

	if err != nil {
		return nil, err
	}

	lock.Lock()

	fonts[path] = v

	lock.Unlock()

	return v, nil
}

func LoadImage(data []byte) (image.Image, string, error) {
	return image.Decode(bytes.NewReader(data))
}
