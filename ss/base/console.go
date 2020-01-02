package base

import (
	"fmt"
	"log"

	"github.com/hailongz/golang/duktape"
	less "github.com/hailongz/golang/serverless/app"
)

func print(ctx *duktape.Context, tag string) {

	n := ctx.GetTop()
	args := []interface{}{tag}

	for i := 0; i < n; i++ {
		switch ctx.GetType(-n + i) {
		case duktape.TypeUndefined:
			args = append(args, "undefined")
			break
		case duktape.TypeNull:
			args = append(args, "null")
			break
		case duktape.TypeBoolean:
			args = append(args, ctx.ToBoolean(-n+i))
			break
		case duktape.TypeNumber:
			args = append(args, ctx.ToNumber(-n+i))
			break
		case duktape.TypeString:
			args = append(args, ctx.ToString(-n+i))
			break
		case duktape.TypeBuffer:
			args = append(args, "buffer")
			break
		case duktape.TypePointer:
			args = append(args, "pointer")
			break
		case duktape.TypeLightFunc:
			args = append(args, "lightfunc")
			break
		case duktape.TypeObject:
			ctx.Dup(-n + i)
			args = append(args, ctx.JsonEncode(-1))
			ctx.Pop()
			break
		}
	}

	log.Println(args...)

}

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		ctx.PushObject()

		ctx.PushGoFunction(func() int {
			print(ctx, fmt.Sprintf("[%s] [INFO]", trace))
			return 0
		})
		ctx.PutPropString(-2, "info")

		ctx.PushGoFunction(func() int {
			print(ctx, fmt.Sprintf("[%s] [DEBUG]", trace))
			return 0
		})
		ctx.PutPropString(-2, "debug")

		ctx.PushGoFunction(func() int {
			print(ctx, fmt.Sprintf("[%s] [ERROR]", trace))
			return 0
		})
		ctx.PutPropString(-2, "error")

		ctx.PushGoFunction(func() int {
			print(ctx, fmt.Sprintf("[%s] [LOG]", trace))
			return 0
		})
		ctx.PutPropString(-2, "log")

		ctx.PushGoFunction(func() int {
			print(ctx, fmt.Sprintf("[%s] [WARN]", trace))
			return 0
		})
		ctx.PutPropString(-2, "warn")

		ctx.PutGlobalString("console")

	})

}
