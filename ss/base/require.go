package base

import (
	"bytes"
	"log"
	"path"
	"strings"

	"github.com/hailongz/golang/duktape"
	less "github.com/hailongz/golang/ss/app"
)

func pushRequire(dir string, ctx *duktape.Context, store less.IStore) {

	ctx.PushGoFunction(func() int {

		n := ctx.GetTop()

		if n > 0 {

			p := ctx.ToString(-n)

			if p != "" {

				if !strings.HasSuffix(p, ".js") {
					p = p + ".js"
				}

				if strings.HasPrefix(p, "/") {
					p = path.Clean("." + p)
				} else {
					p = path.Clean(dir + "/" + p)
				}

				// log.Println(p)

				key := "__m_" + p

				ctx.PushHeapStash()        // 1
				ctx.GetPropString(-1, key) // 2

				if ctx.IsObject(-1) {
					ctx.GetPropString(-1, "exports") // 3
					ctx.Remove(-2)                   // 2
					ctx.Remove(-2)                   // 1
					return 1
				}

				ctx.Pop() // 1

				b, err := store.GetContent(p)

				if err != nil {
					ctx.Pop() // 0
					log.Println("[ERROR]", "[Not Found]", p, err)
					return 0
				}

				ctx.PushObject() // 2

				buf := bytes.NewBuffer(nil)

				buf.WriteString("(function(module,exports,require,__dirname){ ")

				if b != nil {
					buf.Write(b)
				}

				buf.WriteString("})")

				ctx.PushString(p) // 3

				ctx.CompileStringFilename(0, buf.String()) // 3

				if ctx.Pcall(0) != duktape.ExecSuccess {
					err = ctx.ToError(-1)
					ctx.PopN(3) // 0
					log.Println("[ERROR]", err)
					return 0
				}

				ctx.Dup(-2)                      // 4
				ctx.PushObject()                 // 5
				ctx.Dup(-1)                      // 6
				ctx.PutPropString(-3, "exports") // 5
				dir := path.Dir(p)
				pushRequire(dir, ctx, store) // 6
				ctx.PushString(dir)

				if ctx.Pcall(4) != duktape.ExecSuccess {
					err = ctx.ToError(-1)
					ctx.PopN(3)
					log.Println("[ERROR]", err)
					return 0
				} else {
					ctx.Pop()
				}

				ctx.Dup(-1)
				ctx.PutPropString(-3, key)

				ctx.GetPropString(-1, "exports")
				ctx.Remove(-2)
				ctx.Remove(-2)

				return 1
			}
		}

		return 0
	})

}

func init() {

	less.AddOpenlib(func(app *less.App, ctx *duktape.Context, trace string) {

		store := app.GetStore()

		pushRequire("", ctx, store)
		ctx.PutGlobalString("require")
	})

}
