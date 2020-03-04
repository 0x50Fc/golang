package job

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"math/rand"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func Prefix(app micro.IContext, prefix string, jobId int64) string {
	tableCount := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"table", "count"}), 128)
	return fmt.Sprintf("%s%d_", prefix, (jobId%tableCount)+1)
}

func NewToken() string {
	m := md5.New()
	m.Write([]byte(fmt.Sprintf("dkfjwofeiwjf %d dsklfjsdf %d", time.Now().UnixNano(), rand.Int())))
	return hex.EncodeToString(m.Sum(nil))
}
