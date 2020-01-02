package text

import (
	"image"
	"image/color"
	"image/draw"
)

func Draw(dst draw.Image, v *Layout, x, y, width, height float64, pleft float64, pright float64, textAlign TextAlignment) {

	for _, row := range v.Rows {

		dx := pleft

		switch textAlign {
		case TextAlignmentCenter:
			dx = pleft + (float64(width-pleft-pright)-row.Width)*0.5
			break
		case TextAlignmentRight:
			dx = float64(width) - row.Width - pright
			break
		}

		for _, item := range row.Items {

			u := image.NewUniform(&color.RGBA{uint8((item.Style.Color >> 16) & 0x0ff), uint8((item.Style.Color >> 8) & 0x0ff), uint8((item.Style.Color) & 0x0ff), uint8((item.Style.Color >> 24) & 0x0ff)})

			v := v.Text().StringWithRange(item.Loc, item.Len)

			DrawString(item.Style.Font, dst, u, dx+x+item.X, y+row.Y+item.Y, v)

		}
	}
}
