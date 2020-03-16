package geetest

import (
	"time"

	"github.com/GeeTeam/gt3-golang-sdk/geetest"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Validate(app micro.IContext, task *ValidateTask) (interface{}, error) {

	pkey := dynamic.StringValue(dynamic.GetWithKeys(app.GetConfig(), []string{"geetest", task.CaptchaId}), "")

	if pkey == "" {
		return nil, micro.NewError(ERROR_NOT_FOUND, "not found private key")
	}

	redis, prefix, err := app.GetRedis("default")

	if err != nil {
		return nil, err
	}

	s, err := redis.Get(prefix + task.Key).Result()

	if err != nil {
		return nil, err
	}

	geetest := geetest.NewGeetestLib(task.CaptchaId, pkey, 2*time.Second)

	var r bool = false

	if s == "1" {
		r = geetest.SuccessValidate(task.Challenge, task.Validate, task.Seccode, "", "")
	} else {
		r = geetest.FailbackValidate(task.Challenge, task.Validate, task.Seccode)
	}

	redis.Del(prefix + task.Key).Result()

	if !r {
		return nil, micro.NewError(ERROR_NOT_VALIDATE, "validate failed")
	}

	return nil, nil
}
