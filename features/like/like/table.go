package like

import (
	"fmt"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func Prefix(app micro.IContext, prefix string, tid int64) string {
	tableCount := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"table", "count"}), 128)
	return fmt.Sprintf("%s%d_", prefix, (tid%tableCount)+1)
}
