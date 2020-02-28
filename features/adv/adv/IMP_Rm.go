package adv

import (
	"bytes"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (*Adv, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	//查询是否存在
	v := Adv{}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE id = ? AND channel = ? AND position = ?")

	args = append(args, task.Id)
	args = append(args, task.Channel)
	args = append(args, task.Position)

	rs, err := db.Query(conn, &v, prefix, sql.String(), args...)

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
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到广告")
	}

	//删除赞
	_, err = db.Delete(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	return &v, nil
}
