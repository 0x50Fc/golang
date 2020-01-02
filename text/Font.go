package text

import (
	"errors"
	"image"
	"image/draw"
	"strings"
	"sync"
)

type Index interface {
	GetFont() Font
	GetLength() int
}

type Font interface {
	GetSize() float64
	GetFamily() string
	GetAscent() float64
	GetDescent() float64
	Index(vs []rune, i int) Index
	Kern(p, i Index) float64
	Width(i Index) float64
	Draw(i Index, dst draw.Image, u *image.Uniform, x, y, width, height float64) error
}

type vsFont struct {
	vs []Font
}

func (F *vsFont) GetSize() float64 {
	return F.vs[0].GetSize()
}

func (F *vsFont) GetFamily() string {
	return F.vs[0].GetFamily()
}

func (F *vsFont) GetAscent() float64 {
	return F.vs[0].GetAscent()
}

func (F *vsFont) GetDescent() float64 {
	return F.vs[0].GetDescent()
}

func (F *vsFont) Index(vs []rune, idx int) Index {
	for _, f := range F.vs {
		i := f.Index(vs, idx)
		if i != nil {
			return i
		}
	}
	return nil
}

func (F *vsFont) Kern(p, i Index) float64 {
	if p.GetFont() == i.GetFont() {
		return i.GetFont().Kern(p, i)
	}
	return 0
}

func (F *vsFont) Width(i Index) float64 {
	return i.GetFont().Width(i)
}

func (F *vsFont) Draw(i Index, dst draw.Image, u *image.Uniform, x, y, width, height float64) error {
	return i.GetFont().Draw(i, dst, u, x, y, width, height)
}

type FontLibrary func(family string, size float64) (Font, error)
type FontLoad func(path string)

var fontLoadSet = []FontLoad{}

func AddFontLoad(fn FontLoad) {
	fontLoadSet = append(fontLoadSet, fn)
}

func Load(path string) {
	for _, fn := range fontLoadSet {
		fn(path)
	}
}

var lock sync.Mutex
var librarySet = map[string]FontLibrary{}
var librarys = []FontLibrary{}

func GetFamilys() []string {
	vs := []string{}
	lock.Lock()
	for n, _ := range librarySet {
		vs = append(vs, n)
	}
	lock.Unlock()
	return vs
}

func AddFontLibrary(family string, library FontLibrary) {
	lock.Lock()
	defer lock.Unlock()
	_, ok := librarySet[family]
	if !ok {
		librarySet[family] = library
		librarys = append(librarys, library)
	}
}

func NewFont(family string, size float64) (Font, error) {

	vs := []Font{}

	for _, n := range strings.Split(family, ",") {
		lock.Lock()
		v, ok := librarySet[n]
		lock.Unlock()
		if ok {
			f, err := v(family, size)
			if err == nil {
				vs = append(vs, f)
			}
		}
	}

	if len(vs) == 0 {
		lock.Lock()
		for _, v := range librarys {
			f, err := v(family, size)
			if err == nil {
				vs = append(vs, f)
			}
		}
		lock.Unlock()
	}

	if len(vs) == 1 {
		return vs[0], nil
	} else if len(vs) > 0 {
		return &vsFont{vs: vs}, nil
	} else {
		return nil, errors.New("未找到字体")
	}
}

func MeasureString(F Font, v string) float64 {

	var p Index = nil
	var width float64 = 0

	vs := []rune(v)
	n := len(vs)
	idx := 0

	for idx < n {

		i := F.Index(vs, idx)

		if i == nil {
			idx = idx + 1
			continue
		}

		if p != nil {
			width += F.Kern(p, i)
		}
		width += F.Width(i)
		p = i

		idx = idx + i.GetLength()

	}

	return width
}

func DrawString(F Font, dst draw.Image, u *image.Uniform, x, y float64, v string) {

	var p Index = nil
	var width float64 = 0
	var height float64 = 0

	vs := []rune(v)
	n := len(vs)
	idx := 0

	for idx < n {

		i := F.Index(vs, idx)

		if i == nil {
			idx = idx + 1
			continue
		}

		if p != nil {
			x += F.Kern(p, i)
		}

		width = F.Width(i)
		height = i.GetFont().GetDescent() - i.GetFont().GetAscent()

		F.Draw(i, dst, u, x, y, width, height)

		p = i
		x = x + width

		idx = idx + i.GetLength()
	}

}
