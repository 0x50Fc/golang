package redis

import (
	"log"
	"time"

	"github.com/hailongz/golang/duktape"
	"github.com/hailongz/golang/dynamic"
	less "github.com/hailongz/golang/ss/app"
	R "gopkg.in/redis.v5"
)

func getConnWithApp(app *less.App) *R.Client {
	v, _ := app.GetSharedObject("redis", func() (less.SharedObject, error) {
		config := dynamic.Get(app.GetConfig(), "redis")
		addr := dynamic.StringValue(dynamic.Get(config, "addr"), "127.0.0.1:6379")
		password := dynamic.StringValue(dynamic.Get(config, "password"), "")
		db := dynamic.IntValue(dynamic.Get(config, "db"), 0)

		return R.NewClient(&R.Options{
			Addr:     addr,
			Password: password, // no password set
			DB:       int(db),  // use default DB
		}), nil
	})
	return v.(*R.Client)
}

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		config := dynamic.Get(app.GetConfig(), "redis")

		prefix := dynamic.StringValue(dynamic.Get(config, "prefix"), "")

		var conn *R.Client = nil

		getConn := func() *R.Client {
			if conn == nil {
				conn = getConnWithApp(app)
			}
			return conn
		}

		ctx.PushObject()

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 && ctx.IsString(-n) {

				key := prefix + ctx.ToString(-n)

				conn := getConn()

				v, err := conn.Get(key).Result()

				if err != nil {
					log.Printf("[%s] [REDIS] [ERROR] %s\n", trace, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[REDIS] [ERROR] %s", err.Error())
					return duktape.ErrRetError
				}

				ctx.PushString(v)

				return 1
			}

			return 0
		})
		ctx.PutPropString(-2, "get")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 && ctx.IsString(-n) && ctx.IsString(-n+1) {

				key := prefix + ctx.ToString(-n)

				value := ctx.ToString(-n + 1)

				var expires time.Duration = 0

				if n > 2 && ctx.IsNumber(-n+2) {
					expires = time.Duration(ctx.ToInt(-n+2)) * time.Second
				}

				conn := getConn()

				_, err := conn.Set(key, value, expires).Result()

				if err != nil {
					log.Printf("[%s] [REDIS] [ERROR] %s\n", trace, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[REDIS] [ERROR] %s", err.Error())
					return duktape.ErrRetError
				}

			}

			return 0
		})
		ctx.PutPropString(-2, "set")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 && ctx.IsString(-n) {

				key := prefix + ctx.ToString(-n)

				conn := getConn()

				_, err := conn.Del(key).Result()

				if err != nil {
					log.Printf("[%s] [REDIS] [ERROR] %s\n", trace, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[REDIS] [ERROR] %s", err.Error())
					return duktape.ErrRetError
				}

			}

			return 0
		})

		ctx.PutPropString(-2, "del")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 && ctx.IsString(-n) {

				code := ctx.ToString(-n)

				keys := []string{}

				if n > 1 && ctx.IsArray(-n+1) {
					ctx.Enum(-n+1, duktape.DUK_ENUM_ARRAY_INDICES_ONLY)
					for ctx.Next(-1, true) {
						keys = append(keys, ctx.ToString(-1))
						ctx.Pop2()
					}
					ctx.Pop()
				}

				args := []interface{}{}

				if n > 2 && ctx.IsArray(-n+2) {
					ctx.Enum(-n+2, duktape.DUK_ENUM_ARRAY_INDICES_ONLY)
					for ctx.Next(-1, true) {
						args = append(args, less.Decode(ctx, -1))
						ctx.Pop2()
					}
					ctx.Pop()
				}

				conn := getConn()

				rs, err := conn.Eval(code, keys, args...).Result()

				if err != nil {
					log.Printf("[%s] [REDIS] [ERROR] %s\n", trace, err.Error())
					ctx.PushErrorObject(duktape.ErrError, "[REDIS] [ERROR] %s", err.Error())
					return duktape.ErrRetError
				}

				less.Encode(ctx, rs)

				return 1
			}

			return 0
		})

		ctx.PutPropString(-2, "eval")

		ctx.PushString(prefix)
		ctx.PutPropString(-2, "prefix")

		ctx.PutGlobalString("redis")

	})

}
