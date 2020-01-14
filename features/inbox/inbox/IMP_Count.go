package inbox

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Count(app micro.IContext, task *CountTask) (*InboxCountData, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	groupBy := dynamic.StringValue(task.GroupBy, "")

	prefix = Prefix(app, prefix, task.Uid)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	v := Inbox{}

	switch groupBy {
	case GroupBy_mid:
		sql.WriteString(fmt.Sprintf("SELECT COUNT(DISTINCT mid) as n FROM %s", db.TableName(prefix, &v)))
		break
	case GroupBy_fuid:
		sql.WriteString(fmt.Sprintf("SELECT COUNT(DISTINCT fuid) as n FROM %s", db.TableName(prefix, &v)))
		break
	default:
		sql.WriteString(fmt.Sprintf("SELECT COUNT(*) as n FROM %s", db.TableName(prefix, &v)))
		break
	}

	sql.WriteString(" WHERE uid=?")

	args = append(args, task.Uid)

	if task.Fuid != nil {
		sql.WriteString(" AND fuid=?")
		args = append(args, task.Fuid)
	}

	if task.Type != nil {
		sql.WriteString(" AND (type & ?) != 0")
		args = append(args, task.Type)
	}

	if task.Mid != nil {
		sql.WriteString(" AND mid = ?")
		args = append(args, task.Mid)
	}

	if task.Iid != nil {
		sql.WriteString(" AND iid = ?")
		args = append(args, task.Iid)
	}

	if task.TopId != nil {
		sql.WriteString(" AND ctime <= ?")
		args = append(args, task.TopId)
	}

	app.Println("[SQL]", sql.String())

	rs, err := conn.Query(sql.String(), args...)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	var count int64 = 0

	if rs.Next() {
		err = rs.Scan(&count)
		if err != nil {
			return nil, err
		}
	}

	return &InboxCountData{Total: int32(count)}, nil
}
