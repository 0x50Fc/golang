package geetest

import (
	"fmt"
	"time"

	"github.com/GeeTeam/gt3-golang-sdk/geetest"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Reg(app micro.IContext, task *RegTask) (interface{}, error) {

	pkey := dynamic.StringValue(dynamic.GetWithKeys(app.GetConfig(), []string{"geetest", task.CaptchaId}), "")

	if pkey == "" {
		return nil, micro.NewError(ERROR_NOT_FOUND, "not found private key")
	}

	expires := time.Duration(task.Expires) * time.Second

	geetest := geetest.NewGeetestLib(task.CaptchaId, pkey, 2*time.Second)

	s, r := geetest.PreProcess("", "")

	redis, prefix, err := app.GetRedis("default")

	if err != nil {
		return nil, err
	}

	_, err = redis.Set(prefix+task.Key, fmt.Sprintf("%d", s), expires).Result()

	if err != nil {
		return nil, err
	}

	var data interface{} = nil

	err = json.Unmarshal(r, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
