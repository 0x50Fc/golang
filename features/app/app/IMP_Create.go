package app

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/features/id/client"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Create(app micro.IContext, task *CreateTask) (*App, error) {

	cli, err := micro.GetClient(app, "kk-id")

	if err != nil {
		return nil, err
	}

	id, err := client.API_Get(cli, &client.GetTask{})

	if err != nil {
		return nil, err
	}

	conn, p, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix := Prefix(app, p, id)

	v := App{}

	v.Id = id
	v.Title = dynamic.StringValue(task.Title, "")
	v.Ctime = time.Now().Unix()

	if task.Options != nil {

		options := map[string]interface{}{}

		dynamic.Each(v.Options, func(key interface{}, value interface{}) bool {
			options[dynamic.StringValue(key, "")] = value
			return true
		})

		text := dynamic.StringValue(task.Options, "")

		var data interface{} = nil

		json.Unmarshal([]byte(text), &data)

		dynamic.Each(data, func(key interface{}, value interface{}) bool {
			options[dynamic.StringValue(key, "")] = value
			return true
		})

		v.Options = options
	}

	_, err = db.Insert(conn, &v, prefix)

	return &v, nil
}
