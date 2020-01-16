package doc

import (
	"fmt"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func Prefix(app micro.IContext, prefix string, id int64) string {
	tableCount := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"table", "count"}), 128)
	return fmt.Sprintf("%s%d_", prefix, (id%tableCount)+1)
}
