package view

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"strings"

	"golang.org/x/image/font"

	"github.com/hailongz/golang/canvas/gg"
	"github.com/hailongz/golang/duktape"
	"github.com/hailongz/golang/http"
	less "github.com/hailongz/golang/serverless/app"
	"github.com/hailongz/golang/view"
)

type viewProvider struct {
	store    less.IStore
	imageSet map[string]image.Image
}

func (P *viewProvider) GetImage(src string) image.Image {

	v, ok := P.imageSet[src]

	if ok {
		return v
	}

	if strings.HasPrefix(src, "https://") || strings.HasPrefix(src, "http://") {

		options := http.Options{}
		options.Method = "GET"
		options.Url = src
		options.ResponseType = http.OptionResponseTypeByte

		data, err := http.Send(&options)

		if err != nil {
			return nil
		}

		v, _, err = image.Decode(bytes.NewReader(data.([]byte)))

		if err != nil {
			return nil
		}

		P.imageSet[src] = v

		return v
	}

	data, err := P.store.GetContent(src)

	if err != nil {
		return nil
	}

	v, _, err = image.Decode(bytes.NewReader(data))

	if err != nil {
		return nil
	}

	P.imageSet[src] = v

	return v

}

func (P *viewProvider) GetFont(v string) font.Face {
	return nil
}

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		autoId := int(0)

		viewSet := map[int]view.IView{}

		provider := viewProvider{store: app.GetStore(), imageSet: map[string]image.Image{}}

		ctx.PushObject()

		ctx.PushGoFunction(func() int {

			id := autoId + 1

			autoId = id

			v := view.NewView()

			v.SetProvider(&provider)

			viewSet[id] = v

			ctx.PushInt(id)

			return 1
		})

		ctx.PutPropString(-2, "open")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 {

				id := ctx.ToInt(-n)

				delete(viewSet, id)

			}

			return 0
		})

		ctx.PutPropString(-2, "close")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 && ctx.IsString(-n+1) {

				id := ctx.ToInt(-n)
				name := ctx.ToString(-n + 1)

				v, ok := viewSet[id]

				if !ok {
					ctx.PushErrorObject(duktape.ErrError, "[Canvas] [fill] %s", "Not Found Canvas")
					return duktape.ErrRetError
				}

				e := v.CreateElement(name)

				ctx.PushInt(int(e.Id()))

				return 1
			}

			return 0
		})

		ctx.PutPropString(-2, "create")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 2 && ctx.IsString(-n+2) {

				id := ctx.ToInt(-n)
				pid := ctx.ToInt(-n + 1)
				name := ctx.ToString(-n + 2)

				v, ok := viewSet[id]

				if !ok {
					ctx.PushErrorObject(duktape.ErrError, "[View] %s", "Not Found View")
					return duktape.ErrRetError
				}

				p := v.GetElement(int64(pid))

				if p == nil {
					ctx.PushErrorObject(duktape.ErrError, "[View] %s", "Not Found Element")
					return duktape.ErrRetError
				}

				e := v.CreateElement(name)

				ctx.PushString(fmt.Sprintf("%d", e.Id()))

				p.Append(e)

				return 1
			}

			return 0
		})

		ctx.PutPropString(-2, "add")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 {

				id := ctx.ToInt(-n)
				pid := ctx.ToInt(-n + 1)

				v, ok := viewSet[id]

				if !ok {
					ctx.PushErrorObject(duktape.ErrError, "[View] %s", "Not Found View")
					return duktape.ErrRetError
				}

				p := v.GetElement(int64(pid))

				if p == nil {
					ctx.PushErrorObject(duktape.ErrError, "[View] %s", "Not Found Element")
					return duktape.ErrRetError
				}

				p.Remove()
				v.DeleteElement(int64(pid))

			}

			return 0
		})

		ctx.PutPropString(-2, "del")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 2 && ctx.IsString(-n+2) {

				id := ctx.ToInt(-n)
				pid := ctx.ToInt(-n + 1)
				key := ctx.ToString(-n + 2)

				v, ok := viewSet[id]

				if !ok {
					ctx.PushErrorObject(duktape.ErrError, "[View] %s", "Not Found View")
					return duktape.ErrRetError
				}

				var value interface{} = nil

				if n > 3 {
					value = less.Decode(ctx, -n+3)
				}

				v.Set(int64(pid), key, value)

			}

			return 0
		})

		ctx.PutPropString(-2, "set")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 3 {

				id := ctx.ToInt(-n)
				pid := int64(ctx.ToInt(-n + 1))
				width := ctx.ToNumber(-n + 2)
				height := ctx.ToNumber(-n + 3)

				v, ok := viewSet[id]

				if !ok {
					ctx.PushErrorObject(duktape.ErrError, "[View] %s", "Not Found View")
					return duktape.ErrRetError
				}

				w := width
				h := height

				if width < 0 {
					w = math.MaxFloat64
				}

				if height < 0 {
					h = math.MaxFloat64
				}

				v.SetSize(w, h)
				v.Layout(pid)

				if width < 0 || height < 0 {
					width = 0
					height = 0
					e := v.GetElement(pid)
					if e != nil {
						ee, ok := e.(view.IViewElement)
						if ok {
							width = ee.Width()
							height = ee.Height()
						}
					}
				}

				canvas := gg.NewCanvas(int(width), int(height))

				v.Draw(pid, canvas)

				b := bytes.NewBuffer(nil)

				err := png.Encode(b, canvas.Image())

				if err != nil {
					ctx.PushErrorObject(duktape.ErrError, "[Canvas] [drawImage] %s", err.Error())
					return duktape.ErrRetError
				}

				ctx.PushBytes(b.Bytes())

				return 1

			}

			return 0
		})

		ctx.PutPropString(-2, "toPNGData")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 4 {

				id := ctx.ToInt(-n)
				pid := int64(ctx.ToInt(-n + 1))
				width := ctx.ToNumber(-n + 2)
				height := ctx.ToNumber(-n + 3)
				quality := ctx.ToInt(-n + 4)

				v, ok := viewSet[id]

				if !ok {
					ctx.PushErrorObject(duktape.ErrError, "[View] %s", "Not Found View")
					return duktape.ErrRetError
				}

				w := width
				h := height

				if width < 0 {
					w = math.MaxFloat64
				}

				if height < 0 {
					h = math.MaxFloat64
				}

				v.SetSize(w, h)
				v.Layout(pid)

				if width < 0 || height < 0 {
					width = 0
					height = 0
					e := v.GetElement(pid)
					if e != nil {
						ee, ok := e.(view.IViewElement)
						if ok {
							width = ee.Width()
							height = ee.Height()
						}
					}
				}

				canvas := gg.NewCanvas(int(width), int(height))

				v.Draw(pid, canvas)

				b := bytes.NewBuffer(nil)

				if quality < 1 {
					quality = 1
				}

				if quality > 100 {
					quality = 100
				}

				err := jpeg.Encode(b, canvas.Image(), &jpeg.Options{quality})

				if err != nil {
					ctx.PushErrorObject(duktape.ErrError, "[Canvas] [drawImage] %s", err.Error())
					return duktape.ErrRetError
				}

				ctx.PushBytes(b.Bytes())

				return 1

			}

			return 0

		})

		ctx.PutPropString(-2, "toJPGData")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 3 {

				id := ctx.ToInt(-n)
				name := ctx.ToString(-n + 1)
				scale := ctx.ToNumber(-n + 2)
				base := ctx.ToNumber(-n + 3)

				v, ok := viewSet[id]

				if !ok {
					ctx.PushErrorObject(duktape.ErrError, "[View] %s", "Not Found View")
					return duktape.ErrRetError
				}

				v.SetUnit(name, scale, base)

			}

			return 0

		})

		ctx.PutPropString(-2, "setUnit")

		ctx.PutGlobalString("view")

	})

}
