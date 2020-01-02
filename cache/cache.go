package cache

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/hailongz/golang/dynamic"
)

type ICache interface {
	Get(key string, expires time.Duration) (string, error)
	GetItem(key string, iid string, expires time.Duration) (string, error)
	Set(key string, value string, expires time.Duration) error
	SetItem(key string, iid string, value string, expires time.Duration) error
	Del(key string) error
	DelItem(key string, iid string) error
	Expire(key string, expires time.Duration) error
}

func SignKey(args ...interface{}) string {
	m := md5.New()
	s := []byte("_")
	for i, v := range args {
		if i != 0 {
			m.Write(s)
		}
		m.Write([]byte(dynamic.StringValue(v, "")))
	}
	return hex.EncodeToString(m.Sum(nil))
}
