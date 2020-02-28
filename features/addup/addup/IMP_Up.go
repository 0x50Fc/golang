package addup

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Up(app micro.IContext, task *UpTask) (interface{}, error) {

	conn, name, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	if task.Region == nil {
		name = fmt.Sprintf("%s%s", name, task.Name)
	} else {
		name = fmt.Sprintf("%s%d_%s", name, dynamic.IntValue(task.Region, 0), task.Name)
	}

	app.Println("[NAME]", name)

	ret := map[string]interface{}{
		"name": name,
	}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(fmt.Sprintf("SELECT id FROM %s WHERE iid=? AND time=?", name))

	args = append(args, task.Iid, task.Time)

	if task.UnionKeys != nil {

		var unionKeys interface{} = nil

		err = json.Unmarshal([]byte(dynamic.StringValue(task.UnionKeys, "")), &unionKeys)

		if err != nil {
			return nil, err
		}

		dynamic.Each(unionKeys, func(key interface{}, value interface{}) bool {

			sql.WriteString(fmt.Sprintf(" AND `%s`=?", key))
			args = append(args, value)

			return true
		})
	}

	err = db.Transaction(conn, func(conn db.Database) error {

		var id int64 = 0

		rs, err := conn.Query(sql.String(), args...)

		if err != nil {
			return err
		}

		if rs.Next() {
			err = rs.Scan(&id)
		}

		rs.Close()

		if err != nil {
			return err
		}

		sql := bytes.NewBuffer(nil)

		args := []interface{}{}

		if id == 0 {

			sql.WriteString("INSERT INTO ")
			sql.WriteString(name)
			sql.WriteString("(iid,time")

			args = append(args, task.Iid, task.Time)

			n := 0

			if task.Set != nil {
				var data interface{} = nil
				err = json.Unmarshal([]byte(dynamic.StringValue(task.Set, "")), &data)
				if err != nil {
					return err
				}
				dynamic.Each(data, func(key interface{}, value interface{}) bool {
					sql.WriteString(",`")
					sql.WriteString(dynamic.StringValue(key, ""))
					sql.WriteString("`")
					args = append(args, value)
					n = n + 1
					return true
				})
			}

			if task.Add != nil {
				var data interface{} = nil
				err = json.Unmarshal([]byte(dynamic.StringValue(task.Add, "")), &data)
				if err != nil {
					return err
				}
				dynamic.Each(data, func(key interface{}, value interface{}) bool {
					sql.WriteString(",`")
					sql.WriteString(dynamic.StringValue(key, ""))
					sql.WriteString("`")
					args = append(args, value)
					n = n + 1
					return true
				})
			}

			sql.WriteString(") VALUES(?,?")

			for i := 0; i < n; i++ {
				sql.WriteString(",?")
			}

			sql.WriteString(")")

		} else {

			sql.WriteString("UPDATE ")
			sql.WriteString(name)
			sql.WriteString(" SET ")

			n := 0

			if task.Set != nil {
				var data interface{} = nil
				err = json.Unmarshal([]byte(dynamic.StringValue(task.Set, "")), &data)
				if err != nil {
					return err
				}
				dynamic.Each(data, func(key interface{}, value interface{}) bool {
					if n != 0 {
						sql.WriteString(",")
					}
					sql.WriteString("`")
					sql.WriteString(dynamic.StringValue(key, ""))
					sql.WriteString("`=?")
					args = append(args, value)
					n = n + 1
					return true
				})
			}

			if task.Add != nil {
				var data interface{} = nil
				err = json.Unmarshal([]byte(dynamic.StringValue(task.Add, "")), &data)
				if err != nil {
					return err
				}
				dynamic.Each(data, func(key interface{}, value interface{}) bool {
					if n != 0 {
						sql.WriteString(",")
					}
					sql.WriteString("`")
					sql.WriteString(dynamic.StringValue(key, ""))
					sql.WriteString("`=`")
					sql.WriteString(dynamic.StringValue(key, ""))
					sql.WriteString("` + ?")
					args = append(args, value)
					n = n + 1
					return true
				})
			}

			if n == 0 {
				return micro.NewError(ERROR_NOT_FOUND, "无可更新的数据")
			}

			sql.WriteString(" WHERE id=?")

			args = append(args, id)

		}

		r, err := conn.Exec(sql.String(), args...)

		if err != nil {
			return err
		}

		if id == 0 {
			ret["id"], _ = r.LastInsertId()
		} else {
			ret["id"] = id
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	{

		cache, err := app.GetCache("default")
		if err == nil {
			cache.Del(name)
		}
	}

	app.SendMessage(task.GetName(), ret)

	return ret, nil
}
