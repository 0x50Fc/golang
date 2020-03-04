package wx

import (
	"bytes"
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) ContentRm(app micro.IContext, task *ContentRmTask) (interface{}, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	//查询是否存在
	v := Content{}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE id = ?")

	args = append(args, task.Id)

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.Groupid != nil {
		sql.WriteString(" AND groupid=?")
		args = append(args, task.Groupid)
	}

	//先查询是否评论过
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
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到内容")
	}

	//删除赞
	_, err = db.Delete(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
