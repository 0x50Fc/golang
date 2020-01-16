package auth

import (
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (interface{}, error) {

	key := task.Key

	cli, prefix, err := app.GetRedis("default")

	if err != nil {
		return nil, err
	}

	_, err = cli.Del(prefix + key).Result()

	if err != nil {
		return nil, err
	}

	return nil, nil
}
