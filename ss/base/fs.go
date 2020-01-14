package base

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hailongz/golang/duktape"
	"github.com/hailongz/golang/dynamic"
	less "github.com/hailongz/golang/ss/app"
)

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		dir := dynamic.StringValue(dynamic.GetWithKeys(app.GetConfig(), []string{"fs", "dir"}), "")

		ctx.PushObject()

		ctx.PushGoFunction(func() int {

			if dir == "" {
				log.Println("[ERROR] 未找到可用文件目录")
				ctx.PushErrorObject(duktape.ErrError, "未找到可用文件目录", "")
				return duktape.ErrRetError
			}

			n := ctx.GetTop()

			if n > 1 && ctx.IsString(-n) {

				p := ctx.ToString(-n)

				var v []byte = nil

				if ctx.IsBuffer(-n+1) || ctx.IsBufferData(-n+1) {
					v = ctx.ToBytes(-n + 1)
				} else {
					v = []byte(ctx.ToString(-n + 1))
				}

				p = filepath.Clean(filepath.Join(".", p))

				log.Println(p)

				if strings.HasPrefix(p, dir) {
					err := ioutil.WriteFile(p, v, 0777)
					if err != nil {
						log.Printf("[ERROR] 未找到可用文件目录 %s %s\n", p, err.Error())
						ctx.PushErrorObject(duktape.ErrError, "写文件错误 %s", err.Error())
						return duktape.ErrRetError
					}
				} else {
					log.Printf("[ERROR] 路径不可用 %s\n", p)
					ctx.PushErrorObject(duktape.ErrError, "路径不可用 %s", p)
					return duktape.ErrRetError
				}

			}

			return 0
		})

		ctx.PutPropString(-2, "putContent")

		ctx.PushGoFunction(func() int {

			if dir == "" {
				log.Println("[ERROR] 未找到可用文件目录")
				ctx.PushErrorObject(duktape.ErrError, "未找到可用文件目录", "")
				return duktape.ErrRetError
			}

			n := ctx.GetTop()

			if n > 0 && ctx.IsString(-n) {

				p := ctx.ToString(-n)

				p = filepath.Clean(filepath.Join(".", p))

				log.Println(p)

				if strings.HasPrefix(p, dir) {
					err := os.RemoveAll(p)
					if err != nil {
						log.Printf("[ERROR] 未找到可用文件目录 %s %s\n", p, err.Error())
						ctx.PushErrorObject(duktape.ErrError, "写文件错误 %s", err.Error())
						return duktape.ErrRetError
					}
				} else {
					log.Printf("[ERROR] 路径不可用 %s\n", p)
					ctx.PushErrorObject(duktape.ErrError, "路径不可用 %s", p)
					return duktape.ErrRetError
				}

			}

			return 0
		})

		ctx.PutPropString(-2, "remove")

		ctx.PutGlobalString("fs")

	})

}
