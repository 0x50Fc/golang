package vcode

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Create(app micro.IContext, task *CreateTask) (*VCode, error) {

	expires := int64(task.Expires)
	length := int(dynamic.IntValue(task.Length, 4))
	key := task.Key

	if length <= 0 {
		length = 4
	}

	if length > 12 {
		length = 12
	}
	if expires == 0 {
		expires = 60
	}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := VCode{}

	err = func() error {

		rs, err := db.Query(conn, &v, prefix, " WHERE `key`=?", key)

		if err != nil {
			return err
		}

		defer rs.Close()

		if rs.Next() {

			scaner := db.NewScaner(&v)

			err = scaner.Scan(rs)

			if err != nil {
				return err
			}

		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	v.Key = task.Key
	v.Code = NewCode(app, length)
	v.Hash = Hash(v.Code)
	v.Etime = time.Now().Unix() + expires

	if v.Id == 0 {
		_, err = db.Insert(conn, &v, prefix)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = db.UpdateWithKeys(conn, &v, prefix, map[string]bool{"code": true, "hash": true, "etime": true})
		if err != nil {
			return nil, err
		}
	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
