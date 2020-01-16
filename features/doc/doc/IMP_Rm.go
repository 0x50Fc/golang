package doc

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (*Doc, error) {

	v := Doc{}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	sql := bytes.NewBuffer(nil)
	args := []interface{}{}

	sql.WriteString(" WHERE uid=? AND id=?")
	args = append(args, task.Uid, task.Id)

	p, err := db.Get(conn, &v, prefix, sql.String(), args...)

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到文档")
	}

	_, err = db.Delete(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	cache, _ := app.GetCache("default")

	if cache != nil {

		var id int64

		rs, err := conn.Query(fmt.Sprintf("SELECT id FROM %s%s WHERE path LIKE ? ORDER BY id ASC", prefix, v.GetName()), fmt.Sprintf("%s%d/%%", v.Path, v.Id))

		if err != nil {
			return nil, err
		}

		for rs.Next() {

			err = rs.Scan(&id)

			if err != nil {
				rs.Close()
				return nil, err
			}

			cache.Del(fmt.Sprintf("%d/%d", task.Uid, id))
		}

		rs.Close()

	}

	_, err = db.DeleteWithSQL(conn, &v, prefix, " WHERE path LIKE ?", fmt.Sprintf("%s%d/%%", v.Path, v.Id))

	if err != nil {
		return nil, err
	}

	if cache != nil {
		cache.Del(fmt.Sprintf("%d/%d", task.Uid, task.Id))
	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
