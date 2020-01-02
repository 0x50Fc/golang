package user

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hailongz/golang/cache"
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func SetCache(v *User, expires time.Duration, cache cache.ICache) {
	b, _ := json.Marshal(v)
	s := string(b)
	cache.Set(fmt.Sprintf("id_%d", v.Id), s, expires)
	cache.Set(fmt.Sprintf("name_%s", v.Name), s, expires)
	if v.Nick != "" {
		cache.Set(fmt.Sprintf("nick_%s", v.Nick), s, expires)
	}
}

func DelCache(v *User, cache cache.ICache) {
	cache.Del(fmt.Sprintf("id_%d", v.Id))
	cache.Del(fmt.Sprintf("name_%s", v.Name))
	if v.Nick != "" {
		cache.Del(fmt.Sprintf("nick_%s", v.Nick))
	}
}

func (S *Service) Get(app micro.IContext, task *GetTask) (*User, error) {

	if task.Id == nil && task.Name == nil && task.Nick == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到用户ID")
	}

	expires := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 1800)) * time.Second

	cache, _ := app.GetCache("default")

	v := User{}

	sql := bytes.NewBuffer(nil)
	args := []interface{}{}

	if task.Id != nil {
		sql.WriteString("WHERE id=?")
		args = append(args, task.Id)

		if cache != nil {
			s, err := cache.Get(fmt.Sprintf("id_%d", task.Id), expires)
			if err == nil {
				err = json.Unmarshal([]byte(s), &v)
				if err == nil {
					return &v, nil
				}
			}
		}

	} else if task.Name != nil {

		sql.WriteString("WHERE name=?")
		args = append(args, task.Name)

		if cache != nil {
			s, err := cache.Get(fmt.Sprintf("name_%s", task.Name), expires)
			if err == nil {
				err = json.Unmarshal([]byte(s), &v)
				if err == nil {
					return &v, nil
				}
			}
		}
	} else if task.Nick != nil {

		sql.WriteString("WHERE nick=?")
		args = append(args, task.Nick)

		if cache != nil {
			s, err := cache.Get(fmt.Sprintf("nick_%s", task.Nick), expires)
			if err == nil {
				err = json.Unmarshal([]byte(s), &v)
				if err == nil {
					return &v, nil
				}
			}
		}
	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	p, err := db.Get(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	if p == nil {

		if task.Name != nil && dynamic.BooleanValue(task.Autocreate, false) {

			createTask := CreateTask{}
			createTask.Name = dynamic.StringValue(task.Name, "")
			createTask.Nick = task.Nick

			rs, err := S.Create(app, &createTask)

			if err != nil {
				return nil, err
			}

			if cache != nil {
				SetCache(&v, expires, cache)
			}

			return rs, nil
		}

		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到用户")
	}

	if cache != nil {
		SetCache(&v, expires, cache)
	}

	return nil, nil
}
