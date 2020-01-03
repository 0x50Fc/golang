package vcode

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Check(app micro.IContext, task *CheckTask) (*VCode, error) {

	if task.Code == nil && task.Hash == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到验证码")
	}

	key := task.Key

	conn, prefix, err := app.GetDB("rd")

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

		} else {
			return micro.NewError(ERROR_NOT_FOUND, "未找到验证码")
		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	if time.Now().Unix() > v.Etime {
		_, _ = db.Delete(conn, &v, prefix)
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到验证码")
	}

	if task.Code != nil {
		if dynamic.StringValue(task.Code, "") != v.Code {
			return nil, micro.NewError(ERROR_CODE, "错误的验证码")
		}
	} else if task.Hash != nil {
		if dynamic.StringValue(task.Hash, "") != v.Hash {
			return nil, micro.NewError(ERROR_HASH, "错误的验证码")
		}
	}

	return &v, nil
}
