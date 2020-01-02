package redis

import (
	"fmt"
	"log"

	"github.com/hailongz/golang/mq"

	"github.com/hailongz/golang/duktape"
	"github.com/hailongz/golang/dynamic"
	less "github.com/hailongz/golang/serverless/app"
)

func getConnWithApp(app *less.App, name string) (mq.Producer, error) {
	v, err := app.GetSharedObject(fmt.Sprintf("mq.producer.%s", name), func() (less.SharedObject, error) {
		config := dynamic.GetWithKeys(app.GetConfig(), []string{"mq", "producer", name})
		stype := dynamic.StringValue(dynamic.Get(config, "type"), "type")
		return mq.OpenProducer(stype, config)
	})
	if err != nil {
		return nil, err
	}
	return v.(mq.Producer), nil
}

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		var conn mq.Producer = nil

		getConn := func(name string) (mq.Producer, error) {
			var err error = nil
			if conn == nil {
				conn, err = getConnWithApp(app, name)
				if err != nil {
					return nil, err
				}
			}
			return conn, nil
		}

		ctx.PushObject()

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 1 && ctx.IsString(-n) && ctx.IsString(-n+1) {

				conn_name := ctx.ToString(-n)

				name := ctx.ToString(-n + 1)

				conn, err := getConn(conn_name)

				if err != nil {
					ctx.PushErrorObject(duktape.ErrError, "[MQ] [ERROR] [OPEN] %s", err.Error())
					log.Printf("[MQ] [ERROR] [OPEN] %s\n", err.Error())
					return duktape.ErrRetError
				}

				var data interface{} = nil

				if n > 2 {
					data = less.Decode(ctx, -n+2)
				}

				err = conn.Send(name, data)

				if err != nil {
					ctx.PushErrorObject(duktape.ErrError, "[MQ] [ERROR] [SEND] %s", err.Error())
					log.Printf("[MQ] [ERROR] [SEND] %s\n", err.Error())
					return duktape.ErrRetError
				}
			}

			return 0
		})

		ctx.PutPropString(-2, "send")

		ctx.PutGlobalString("mq")

	})

}
