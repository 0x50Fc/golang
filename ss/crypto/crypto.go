package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"

	"github.com/hailongz/golang/duktape"
	less "github.com/hailongz/golang/ss/app"
)

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		ctx.PushObject()

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 && ctx.IsString(-n) {

				s := ctx.ToString(-n)

				m := md5.New()
				m.Write([]byte(s))

				v := hex.EncodeToString(m.Sum(nil))

				ctx.PushString(v)

				return 1
			}

			return 0
		})
		ctx.PutPropString(-2, "md5")

		ctx.PushGoFunction(func() int {

			n := ctx.GetTop()

			if n > 0 && ctx.IsString(-n) {

				s := ctx.ToString(-n)

				m := sha1.New()
				m.Write([]byte(s))

				v := hex.EncodeToString(m.Sum(nil))

				ctx.PushString(v)

				return 1
			}

			return 0
		})
		ctx.PutPropString(-2, "sha1")

		ctx.PutGlobalString("crypto")

	})

}
