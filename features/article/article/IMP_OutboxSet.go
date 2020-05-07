package article

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) OutboxSet(app micro.IContext, task *OutboxSetTask) (*Outbox, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	v := Outbox{}

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
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到草稿")
	}

	keys := map[string]bool{}

	if task.Body != nil {
		v.Body = dynamic.StringValue(task.Body, "")
		keys["body"] = true
	}

	var options interface{} = nil

	if task.Options != nil {

		text := dynamic.StringValue(task.Options, "")
		var data interface{} = nil
		json.Unmarshal([]byte(text), &data)

		options = db.Merge(v.Options, data)

		v.Options = data
		keys["options"] = true
	}

	if len(keys) > 0 {

		_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

		if err != nil {
			return nil, err
		}

		if keys["options"] {
			v.Options = options
		}

		app.SendMessage(task.GetName(), &v)
	}

	return &v, nil
}
