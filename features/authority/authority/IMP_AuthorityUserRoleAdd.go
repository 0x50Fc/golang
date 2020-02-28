package authority

import (
	"encoding/json"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AuthorityUserRoleAdd(app micro.IContext, task *AuthorityUserRoleAddTask) (*Authority, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Authority{}

	err = db.Transaction(conn, func(conn db.Database) error {

		_, err := db.Get(conn, &v, prefix, " WHERE uid=? AND roleid=? AND resid=0", task.Uid, task.RoleId)

		if err != nil {
			return err
		}

		keys := map[string]bool{}

		v.Uid = task.Uid
		v.RoleId = task.RoleId

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

		if v.Id == 0 {

			_, err = db.Insert(conn, &v, prefix)

			if err != nil {
				return err
			}

		} else if len(keys) > 0 {

			_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

			if err != nil {
				return err
			}

		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &v, nil
}
