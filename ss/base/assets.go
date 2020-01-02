package base

import (
	"log"

	"github.com/hailongz/golang/duktape"
	less "github.com/hailongz/golang/ss/app"
)

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		store := app.GetStore()

		ctx.PushObject()

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 && ctx.IsString(-n) {

				p := ctx.ToString(-n)

				b, err := store.GetContent(p)

				if err != nil {
					log.Println("[ERROR]", "[Not Found]", p, err)
				}

				if b != nil {
					ctx.PushString(string(b))
					return 1
				}

			}

			return 0
		})

		ctx.PutPropString(-2, "getString")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 && ctx.IsString(-n) {

				p := ctx.ToString(-n)

				b, err := store.GetContent(p)

				if err != nil {
					log.Println("[ERROR]", "[Not Found]", p, err)
				}

				if b != nil {
					ctx.PushBytes(b)
					return 1
				}

			}

			return 0
		})

		ctx.PutPropString(-2, "get")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 && ctx.IsString(-n) {

				p := ctx.ToString(-n)

				b, mt := store.Has(p)

				if b {
					ctx.PushInt(int(mt.Unix()))
					return 1
				}

			}

			return 0
		})

		ctx.PutPropString(-2, "has")

		ctx.PutGlobalString("assets")

	})

}
