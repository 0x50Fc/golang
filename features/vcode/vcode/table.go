package vcode

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"

	"github.com/hailongz/golang/micro"
)

func NewCode(app micro.IContext, length int) string {
	v := fmt.Sprintf("%d", rand.Int())
	for len(v) < length {
		v = fmt.Sprintf("%s%d", v, rand.Int())
	}
	return v[0:length]
}

func Hash(code string) string {
	m := md5.New()
	m.Write([]byte(fmt.Sprintf("*&^YTGBNM<L:P1kedmfsf,%s", code)))
	return hex.EncodeToString(m.Sum(nil))
}
