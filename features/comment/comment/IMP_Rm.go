package comment

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (*Comment, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Eid)

	//查询是否存在
	v := Comment{}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE id = ? AND eid = ?")

	args = append(args, task.Id)
	args = append(args, task.Eid)

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
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
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到评论")
	}

	//删除赞
	_, err = db.Delete(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	//清除缓存
	{
		cache, err := app.GetCache("default")

		if err == nil {

			cache.Del(fmt.Sprintf("%d", task.Eid))

		}
	}

	//MQ
	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
