package text

import (
	"errors"
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
	"golang.org/x/image/vector"
)

func init() {

	AddFontLoad(func(path string) {
		if strings.HasSuffix(path, ".ttf") || strings.HasSuffix(path, ".otf") {

			fd, err := os.Open(path)
			if err != nil {
				log.Println("[Font]", err)
				return
			}
			defer fd.Close()
			bb, err := ioutil.ReadAll(fd)
			if err != nil {
				log.Println("[Font]", err)
				return
			}
			f, err := sfnt.Parse(bb)
			if err != nil {
				log.Println("[Font]", err)
				return
			}
			var b sfnt.Buffer
			name, err := f.Name(&b, sfnt.NameIDFamily)
			if err != nil {
				log.Println("[Font]", err)
				return
			}

			AddFontLibrary(name, func(family string, size float64) (Font, error) {
				var b sfnt.Buffer
				r, err := f.Bounds(&b, fixed.Int26_6(size*64), font.HintingNone)

				if err != nil {
					return nil, err
				}

				return &sFont{family: family, size: size, sf: f, ascent: -float64(r.Max.Y) / 64, descent: -float64(r.Min.Y) / 64, isize: fixed.Int26_6(size * 64)}, nil

			})

			log.Println("[Font] [Load]", path, name)

		}
	})

}

type sIndex struct {
	font  *sFont
	index sfnt.GlyphIndex
}

func (I *sIndex) GetFont() Font {
	return I.font
}

func (I *sIndex) GetLength() int {
	return 1
}

type sFont struct {
	size    float64
	family  string
	sf      *sfnt.Font
	ascent  float64
	descent float64
	isize   fixed.Int26_6
}

func (F *sFont) GetSize() float64 {
	return F.size
}

func (F *sFont) GetFamily() string {
	return F.family
}

func (F *sFont) GetAscent() float64 {
	return F.ascent
}

func (F *sFont) GetDescent() float64 {
	return F.descent
}

func (F *sFont) Index(vs []rune, idx int) Index {
	var b sfnt.Buffer
	i, err := F.sf.GlyphIndex(&b, vs[idx])
	if err == nil {
		return &sIndex{font: F, index: i}
	}
	return nil
}

func (F *sFont) Kern(p, i Index) float64 {
	if p.GetFont() == i.GetFont() {
		pi, pok := p.(*sIndex)
		si, ok := i.(*sIndex)
		if ok && pok {
			var b sfnt.Buffer
			r, err := si.font.sf.Kern(&b, sfnt.GlyphIndex(pi.index), sfnt.GlyphIndex(si.index), si.font.isize, font.HintingNone)
			if err == nil {
				return float64(r) / 64
			}
		}
	}
	return 0
}

func (F *sFont) Width(i Index) float64 {
	var b sfnt.Buffer
	si, ok := i.(*sIndex)
	if ok {
		r, err := si.font.sf.GlyphAdvance(&b, sfnt.GlyphIndex(si.index), si.font.isize, font.HintingNone)
		if err == nil {
			return float64(r) / 64
		}
	}
	return 0
}

func (F *sFont) Draw(i Index, dst draw.Image, u *image.Uniform, x, y, width, height float64) error {

	si, ok := i.(*sIndex)

	if !ok {
		return errors.New("不支持的字体绘制")
	}

	var b sfnt.Buffer

	segments, err := F.sf.LoadGlyph(&b, sfnt.GlyphIndex(si.index), si.font.isize, nil)

	if err != nil {
		return err
	}

	dx := int(x)
	dy := int(y)
	dw := int(math.Ceil(width))
	dh := int(math.Ceil(height))

	iy := float32(F.descent)

	r := vector.NewRasterizer(dw, dh)

	r.DrawOp = draw.Src

	for _, seg := range segments {

		switch seg.Op {
		case sfnt.SegmentOpMoveTo:
			r.MoveTo(
				float32(seg.Args[0].X)/64,
				iy+float32(seg.Args[0].Y)/64,
			)
		case sfnt.SegmentOpLineTo:
			r.LineTo(
				float32(seg.Args[0].X)/64,
				iy+float32(seg.Args[0].Y)/64,
			)
		case sfnt.SegmentOpQuadTo:
			r.QuadTo(
				float32(seg.Args[0].X)/64,
				iy+float32(seg.Args[0].Y)/64,
				float32(seg.Args[1].X)/64,
				iy+float32(seg.Args[1].Y)/64,
			)
		case sfnt.SegmentOpCubeTo:
			r.CubeTo(
				float32(seg.Args[0].X)/64,
				iy+float32(seg.Args[0].Y)/64,
				float32(seg.Args[1].X)/64,
				iy+float32(seg.Args[1].Y)/64,
				float32(seg.Args[2].X)/64,
				iy+float32(seg.Args[2].Y)/64,
			)
		}
	}

	r.Draw(dst, image.Rect(dx, dy, dx+dw, dy+dh), u, image.Pt(0, 0))

	return nil
}
