package client

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func API_Get(cli micro.Client, task *GetTask) (int64, error) {

	var ret int64 = 0

	data, err := cli.Send("GET", task.GetName(), task)

	if err != nil {
		return 0, err
	}

	dynamic.SetValue(&ret, data)

	return ret, nil
}
