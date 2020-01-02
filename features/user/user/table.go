package user

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func Prefix(app micro.IContext, prefix string, uid int64) string {
	tableCount := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"table", "count"}), 128)
	return fmt.Sprintf("%s%d_", prefix, (uid%tableCount)+1)
}

func NewPassword() string {
	return fmt.Sprintf("JIEOOL9392uefjnlJKLJD:OF(*YHJNM%d%d", time.Now().UnixNano(), rand.Int())
}

func EncPassword(v string) string {
	m := md5.New()
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum(nil))
}
