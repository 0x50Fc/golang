package cache

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gopkg.in/redis.v5"
)

type redisCache struct {
	prefix string
	cli    *redis.Client
}

func NewRedisCache(cli *redis.Client, prefix string) ICache {

	v := redisCache{
		cli:    cli,
		prefix: prefix,
	}

	return &v
}

func (R *redisCache) Get(key string, expires time.Duration) (string, error) {
	text, err := R.cli.Get(R.prefix + key).Result()
	if err != nil {
		return text, err
	}
	if strings.HasPrefix(text, "etime:") {
		i := strings.Index(text, ",")
		if i > 0 {

			etime, _ := strconv.ParseInt(text[6:i], 10, 64)

			if time.Now().Unix() > etime {
				R.cli.Expire(key, expires).Result()
			}

			s := text[i+1:]

			return s, nil
		}
	}
	return text, nil
}

func (R *redisCache) GetItem(key string, iid string, expires time.Duration) (string, error) {

	text, err := R.cli.HGet(R.prefix+key, iid).Result()

	if err != nil {
		return text, err
	}

	if strings.HasPrefix(text, "etime:") {

		i := strings.Index(text, ",")

		if i > 0 {

			etime, _ := strconv.ParseInt(text[6:i], 10, 64)

			if time.Now().Unix() > etime {
				R.cli.Expire(key, expires).Result()
			}

			s := text[i+1:]

			return s, nil
		}
	}
	return text, nil
}

func (R *redisCache) Set(key string, value string, expires time.Duration) error {
	_, err := R.cli.Set(R.prefix+key, fmt.Sprintf("etime:%d,%s", time.Now().Add(expires-60*time.Second).Unix(), value), expires).Result()
	return err
}

func (R *redisCache) SetItem(key string, iid string, value string, expires time.Duration) error {
	_, err := R.cli.HSet(R.prefix+key, iid, fmt.Sprintf("etime:%d,%s", time.Now().Add(expires-60*time.Second).Unix(), value)).Result()
	if err == nil {
		_, err = R.cli.Expire(R.prefix+key, expires).Result()
	}
	return err
}

func (R *redisCache) Del(key string) error {
	_, err := R.cli.Del(R.prefix + key).Result()
	return err
}

func (R *redisCache) DelItem(key string, iid string) error {
	_, err := R.cli.HDel(R.prefix+key, iid).Result()
	return err
}

func (R *redisCache) Expire(key string, expires time.Duration) error {
	_, err := R.cli.Expire(R.prefix+key, expires).Result()
	return err
}
