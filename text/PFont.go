package text

import (
	"errors"
	"image"
	"image/draw"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/hailongz/golang/emoji"
	"github.com/nfnt/resize"
)

func init() {

	AddFontLoad(func(path string) {

		if strings.HasSuffix(path, ".png") {
			info, err := os.Stat(path)
			if err == nil && info.IsDir() {
				name := filepath.Base(path)
				ext := filepath.Ext(path)
				if ext != "" {
					name = name[:len(name)-len(ext)]
				}
				name = strings.ReplaceAll(name, "-", " ")

				AddFontLibrary(name, func(family string, size float64) (Font, error) {
					p, _ := filepath.Abs(path)
					v := &pFont{basePath: filepath.Clean(p), size: size, family: family, descent: size, ascent: 0, e: emoji.NewEmoji()}
					err := v.emoji()
					if err != nil {
						return nil, err
					}
					return v, nil
				})

			}
		}

	})
}

type pIndex struct {
	font *pFont
	src  image.Image
	n    int
}

func (I *pIndex) GetFont() Font {
	return I.font
}

func (I *pIndex) GetLength() int {
	return I.n
}

type pFont struct {
	size     float64
	family   string
	ascent   float64
	descent  float64
	basePath string
	e        *emoji.Emoji
}

func (F *pFont) emoji() error {

	ext := filepath.Ext(F.basePath)

	return filepath.Walk(F.basePath, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() && strings.HasSuffix(path, ext) {
			name := filepath.Base(path)
			name = name[:len(name)-len(ext)]
			vs := []rune{}
			for _, s := range strings.Split(name, "-") {
				v, _ := strconv.ParseInt(s, 16, 32)
				vs = append(vs, rune(v))
			}
			F.e.Add(vs, path)
		}

		return nil
	})
}

func (F *pFont) GetSize() float64 {
	return F.size
}

func (F *pFont) GetFamily() string {
	return F.family
}

func (F *pFont) GetAscent() float64 {
	return F.ascent
}

func (F *pFont) GetDescent() float64 {
	return F.descent
}

func (F *pFont) Index(vs []rune, idx int) Index {

	i := F.e.Index(vs, idx)

	if i != nil {
		src, err := i.GetImage()
		if err == nil && src != nil {
			return &pIndex{font: F, src: src, n: len(i.GetRune())}
		}
	}

	return nil
}

func (F *pFont) Kern(p, i Index) float64 {
	return 0
}

func (F *pFont) Width(i Index) float64 {
	return F.size
}

func (F *pFont) Draw(i Index, dst draw.Image, u *image.Uniform, x, y, width, height float64) error {

	pi, ok := i.(*pIndex)

	if !ok {
		return errors.New("不支持的字体绘制")
	}

	if pi.src != nil {

		dx := int(x)
		dy := int(y)
		dw := int(math.Ceil(width))
		dh := int(math.Ceil(height))
		r := image.Rect(dx, dy, dx+dw, dy+dh)
		draw.Draw(dst, r, resize.Resize(uint(dw), uint(dh), pi.src, resize.Lanczos3), image.Pt(0, 0), draw.Src)
	}

	return nil
}
