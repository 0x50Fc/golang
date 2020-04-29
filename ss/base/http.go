package base

import (
	"log"
	"time"

	"github.com/hailongz/golang/duktape"
	"github.com/hailongz/golang/http"
	less "github.com/hailongz/golang/ss/app"
)

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		ctx.PushObject()

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 && ctx.IsObject(-n) {

				options := http.Options{}

				{
					ctx.GetPropString(-n, "url")
					if ctx.IsString(-1) {
						options.Url = ctx.ToString(-1)
					}
					ctx.Pop()
				}

				{
					ctx.GetPropString(-n, "method")
					if ctx.IsString(-1) {
						options.Method = ctx.ToString(-1)
					}
					ctx.Pop()
				}

				{
					ctx.GetPropString(-n, "type")
					if ctx.IsString(-1) {
						options.Type = ctx.ToString(-1)
					}
					ctx.Pop()
				}

				{
					ctx.GetPropString(-n, "timeout")
					if ctx.IsNumber(-1) {
						options.Timeout = time.Duration(ctx.ToInt(-1)) * time.Millisecond
					}
					ctx.Pop()
				}

				{
					ctx.GetPropString(-n, "responseType")
					if ctx.IsString(-1) {
						options.ResponseType = ctx.ToString(-1)
					}
					ctx.Pop()
				}

				{
					ctx.GetPropString(-n, "charset")
					if ctx.IsString(-1) {
						options.ResponseCharset = ctx.ToString(-1)
					}
					ctx.Pop()
				}

				{
					ctx.GetPropString(-n, "data")
					less.Unmarshal(ctx, -1, &options.Data)
					ctx.Pop()
				}

				{
					ctx.GetPropString(-n, "header")
					options.Headers = map[string]string{}

					if ctx.IsObject(-1) {

						ctx.Enum(-1, duktape.DUK_ENUM_INCLUDE_SYMBOLS)

						for ctx.Next(-1, true) {
							key := ctx.ToString(-2)
							value := ctx.ToString(-1)

							ctx.Pop2()

							options.Headers[key] = value
						}

						ctx.Pop()
					}

					ctx.Pop()
				}

				options.Headers["Trace-ID"] = trace

				clientIp := ""
				sessionId := ""

				{
					ctx.GetGlobalString("input")
					ctx.GetPropString(-1, "clientIp")
					clientIp = ctx.ToString(-1)
					ctx.Pop()
					ctx.GetPropString(-1, "sessionId")
					sessionId = ctx.ToString(-1)
					ctx.Pop()
					ctx.Pop()
				}

				tv_in := time.Now().UnixNano() / int64(time.Millisecond)

				defer func() {
					now := time.Now()
					tv_out := now.UnixNano() / int64(time.Millisecond)
					log.Printf("[STAT] [%s] [%s] [%s] [in:%d] [out:%d] [use:%d] [clientIp:%s] [session:%s]\n", trace, app.QName(), options.Url, tv_in, tv_out, tv_out-tv_in, clientIp, sessionId)
				}()

				rs, err := http.Send(&options)

				if err != nil {
					ctx.PushErrorObject(duktape.ErrError, "[HTTP] [ERROR] %s", err.Error())
					return duktape.ErrRetError
				}

				less.Encode(ctx, rs)

				return 1
			}

			return 0
		})

		ctx.PutPropString(-2, "send")

		ctx.PutGlobalString("http")

		{
			ctx.PushObject()

			ctx.PushString(http.OptionTypeUrlencode)
			ctx.PutPropString(-2, "Urlencode")

			ctx.PushString(http.OptionTypeJson)
			ctx.PutPropString(-2, "Json")

			ctx.PushString(http.OptionTypeText)
			ctx.PutPropString(-2, "Text")

			ctx.PushString(http.OptionTypeXml)
			ctx.PutPropString(-2, "Xml")

			ctx.PushString(http.OptionTypeMultipart)
			ctx.PutPropString(-2, "FormData")

			ctx.PutGlobalString("HttpType")
		}

		{
			ctx.PushObject()

			ctx.PushString(http.OptionResponseTypeText)
			ctx.PutPropString(-2, "Text")

			ctx.PushString(http.OptionResponseTypeJson)
			ctx.PutPropString(-2, "Json")

			ctx.PushString(http.OptionResponseTypeByte)
			ctx.PutPropString(-2, "Byte")

			ctx.PushString("response")
			ctx.PutPropString(-2, "Response")

			ctx.PutGlobalString("HttpResponseType")
		}

	})

}
