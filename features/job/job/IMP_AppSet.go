package job

import (
	"encoding/json"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AppSet(app micro.IContext, task *AppSetTask) (*App, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := App{}

	rs, err := db.Query(conn, &v, prefix, " WHERE id=?", task.Id)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	if rs.Next() {

		scaner := db.NewScaner(&v)

		err = scaner.Scan(rs)

		if err != nil {
			return nil, err
		}

	} else {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到分组")
	}

	keys := map[string]bool{}

	if task.Alias != nil {
		v.Alias = dynamic.StringValue(task.Alias, "")
		keys["alias"] = true
	}

	if task.Type != nil {
		v.Type = dynamic.StringValue(task.Type, "")
		keys["type"] = true
	}

	if task.Content != nil {
		v.Content = dynamic.StringValue(task.Content, "")
		keys["content"] = true
	}

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
		keys["options"] = true
	}

	if len(keys) > 0 {

		_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

		if err != nil {
			return nil, err
		}

		app.SendMessage(task.GetName(), &v)
	}

	return &v, nil
}
