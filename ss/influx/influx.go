package influx

import (
	"strconv"
	"time"

	"github.com/hailongz/golang/duktape"
	"github.com/hailongz/golang/dynamic"
	less "github.com/hailongz/golang/ss/app"
	influx "github.com/influxdata/influxdb1-client/v2"
)

func newPoint(ctx *duktape.Context, idx int) *influx.Point {
	if ctx.IsObject(idx) {

		name := ""

		ctx.GetPropString(idx, "name")

		if ctx.IsString(-1) {
			name = ctx.ToString(-1)
		}

		ctx.Pop()

		if name == "" {
			return nil
		}

		tags := map[string]string{}
		fields := map[string]interface{}{}

		ctx.GetPropString(idx, "tags")

		if ctx.IsObject(-1) {
			ctx.Enum(-1, duktape.DUK_ENUM_INCLUDE_SYMBOLS)
			for ctx.Next(-1, true) {
				key := ctx.ToString(-2)
				value := ctx.ToString(-1)
				tags[key] = value
				ctx.Pop2()
			}
			ctx.Pop()
		}

		ctx.Pop()

		ctx.GetPropString(idx, "fields")

		if ctx.IsObject(-1) {
			ctx.Enum(-1, duktape.DUK_ENUM_INCLUDE_SYMBOLS)
			for ctx.Next(-1, true) {
				key := ctx.ToString(-2)
				value := less.Decode(ctx, -1)
				fields[key] = value
				ctx.Pop2()
			}
			ctx.Pop()
		}

		ctx.Pop()

		ctx.GetPropString(idx, "time")

		millisecond, _ := strconv.ParseInt(ctx.ToString(-1), 10, 64)

		tv := time.Unix(millisecond/1000, (millisecond%1000)*int64(time.Millisecond))

		ctx.Pop()

		p, err := influx.NewPoint(name, tags, fields, tv)

		if err != nil {
			return nil
		}

		return p
	}
	return nil
}

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		config := dynamic.Get(app.GetConfig(), "influx")

		db := dynamic.StringValue(dynamic.Get(config, "db"), "")

		var conn influx.Client = nil

		getConn := func() (influx.Client, error) {

			if conn == nil {

				v, err := app.GetSharedObject("influx", func() (less.SharedObject, error) {

					addr := dynamic.StringValue(dynamic.Get(config, "addr"), "")
					user := dynamic.StringValue(dynamic.Get(config, "user"), "")
					password := dynamic.StringValue(dynamic.Get(config, "password"), "")

					v, err := influx.NewHTTPClient(influx.HTTPConfig{
						Addr:     addr,
						Username: user,
						Password: password,
					})

					if err != nil {
						return nil, err
					}

					return v, nil
				})

				if err != nil {
					return nil, err
				}

				conn = v.(influx.Client)

			}

			return conn, nil
		}

		ctx.PushObject()

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 {

				bp, _ := influx.NewBatchPoints(influx.BatchPointsConfig{
					Database:  db,
					Precision: "us",
				})

				if ctx.IsArray(-n) {

					ctx.Enum(-n, duktape.DUK_ENUM_ARRAY_INDICES_ONLY)

					for ctx.Next(-1, true) {

						p := newPoint(ctx, -1)

						if p != nil {
							bp.AddPoint(p)
						}

						ctx.Pop2()
					}

					ctx.Pop()

				} else if ctx.IsObject(-n) {

					p := newPoint(ctx, -1)

					if p != nil {
						bp.AddPoint(p)
					}

				} else {
					return 0
				}

				n := len(bp.Points())

				if n > 0 {

					conn, err := getConn()

					if err != nil {
						ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
						return duktape.ErrRetError
					}

					err = conn.Write(bp)

					if err != nil {
						ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
						return duktape.ErrRetError
					}

					ctx.PushInt(n)

					return 1
				}

			}

			return 0
		})
		ctx.PutPropString(-2, "write")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 && ctx.IsString(-n) {

				sql := ctx.ToString(-n)

				conn, err := getConn()

				if err != nil {
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				q := influx.Query{
					Command:  sql,
					Database: db,
				}

				resp, err := conn.Query(q)

				if err == nil {
					err = resp.Error()
				}

				if err != nil {
					ctx.PushErrorObject(duktape.ErrError, "[DB] %s", err.Error())
					return duktape.ErrRetError
				}

				if len(resp.Results) > 0 {

					less.Encode(ctx, resp.Results[0].Series)

					return 1

				}

			}

			return 0
		})

		ctx.PutPropString(-2, "query")

		ctx.PutGlobalString("influx")
	})
}
