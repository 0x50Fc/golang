package base

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hailongz/golang/duktape"
	less "github.com/hailongz/golang/ss/app"
)

func toInt64(ctx *duktape.Context, idx int, defaultValue int64) int64 {
	switch ctx.GetType(idx) {
	case duktape.TypeNumber:
		return int64(ctx.ToNumber(idx))
	default:
		v := ctx.ToString(idx)
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return defaultValue
		}
		return i
	}
}

func toFloat64(ctx *duktape.Context, idx int, defaultValue float64) float64 {
	switch ctx.GetType(idx) {
	case duktape.TypeNumber:
		return ctx.ToNumber(idx)
	default:
		v := ctx.ToString(idx)
		i, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return defaultValue
		}
		return i
	}
}

func isFloat64(ctx *duktape.Context, idx int) bool {
	switch ctx.GetType(idx) {
	case duktape.TypeNumber:
		v := ctx.ToNumber(idx)
		iv := int64(v)
		return v != float64(iv)
	default:
		v := ctx.ToString(idx)
		return strings.Contains(v, ".")
	}
}

func pushInt64(ctx *duktape.Context, v int64) {
	i := float64(v)
	if v == int64(i) {
		ctx.PushNumber(i)
	} else {
		ctx.PushString(fmt.Sprintf("%d", v))
	}
}

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		ctx.PushObject()

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 {
				if isFloat64(ctx, -n) || isFloat64(ctx, -n+1) {
					a := toFloat64(ctx, -n, 0)
					b := toFloat64(ctx, -n+1, 0)
					pushInt64(ctx, int64(a+b))
				} else {
					a := toInt64(ctx, -n, 0)
					b := toInt64(ctx, -n+1, 0)
					pushInt64(ctx, a+b)
				}
				return 1
			}

			return 0
		})
		ctx.PutPropString(-2, "add")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 {
				if isFloat64(ctx, -n) || isFloat64(ctx, -n+1) {
					a := toFloat64(ctx, -n, 0)
					b := toFloat64(ctx, -n+1, 0)
					pushInt64(ctx, int64(a-b))
				} else {
					a := toInt64(ctx, -n, 0)
					b := toInt64(ctx, -n+1, 0)
					pushInt64(ctx, a-b)
				}
				return 1
			}

			return 0
		})
		ctx.PutPropString(-2, "sub")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 {
				if isFloat64(ctx, -n) || isFloat64(ctx, -n+1) {
					a := toFloat64(ctx, -n, 0)
					b := toFloat64(ctx, -n+1, 0)
					pushInt64(ctx, int64(a*b))
				} else {
					a := toInt64(ctx, -n, 0)
					b := toInt64(ctx, -n+1, 0)
					pushInt64(ctx, a*b)
				}

				return 1
			}

			return 0

		})
		ctx.PutPropString(-2, "mul")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 {
				if isFloat64(ctx, -n) || isFloat64(ctx, -n+1) {
					a := toFloat64(ctx, -n, 0)
					b := toFloat64(ctx, -n+1, 0)
					if b == 0 {
						return 0
					}
					pushInt64(ctx, int64(a/b))
				} else {
					a := toInt64(ctx, -n, 0)
					b := toInt64(ctx, -n+1, 0)
					if b == 0 {
						return 0
					}
					pushInt64(ctx, a/b)
				}
				return 1
			}

			return 0
		})
		ctx.PutPropString(-2, "div")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 {
				a := toInt64(ctx, -n, 0)
				b := toInt64(ctx, -n+1, 0)
				pushInt64(ctx, a%b)
				return 1
			}

			return 0
		})

		ctx.PutPropString(-2, "mod")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 {

				if isFloat64(ctx, -n) || isFloat64(ctx, -n+1) {
					a := toFloat64(ctx, -n, 0)
					b := toFloat64(ctx, -n+1, 0)
					if a == b {
						ctx.PushInt(0)
					} else if a > b {
						ctx.PushInt(1)
					} else {
						ctx.PushInt(-1)
					}
				} else {
					a := toInt64(ctx, -n, 0)
					b := toInt64(ctx, -n+1, 0)
					if a == b {
						ctx.PushInt(0)
					} else if a > b {
						ctx.PushInt(1)
					} else {
						ctx.PushInt(-1)
					}
				}

				return 1
			}

			return 0
		})

		ctx.PutPropString(-2, "comp")

		ctx.PutGlobalString("Int64")

	})

}
