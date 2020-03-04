package wx

import (
	"bytes"
	"fmt"
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) ContentCount(app micro.IContext, task *ContentCountTask) (*CountData, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE 1")

	if task.Id != nil {
		sql.WriteString(" AND id = ?")
		args = append(args, task.Id)
	}

	if task.MsgId != "" {
		sql.WriteString(" AND msgId = ?")
		args = append(args, task.MsgId)
	}

	if task.CreateTime != "" {
		sql.WriteString(" AND createTime = ?")
		args = append(args, task.CreateTime)
	}

	if task.MsgTalkerUserAlias != "" {
		sql.WriteString(" AND msgTalkerUserAlias = ?")
		args = append(args, task.MsgTalkerUserAlias)
	}

	if task.MsgTalkerUserNickName != "" {
		sql.WriteString(" AND msgTalkerUserNickName LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.MsgTalkerUserNickName))
	}

	if task.Type != "" {
		sql.WriteString(" AND type = ?")
		args = append(args, task.Type)
	}

	if task.MsgGroupName != "" {
		sql.WriteString(" AND msgGroupName LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.MsgGroupName))
	}

	if task.Q != nil {
		sql.WriteString(" AND msgcontent LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.Q))
	}

	if task.StartTime != nil {
		sql.WriteString(" AND ctime>=?")
		args = append(args, task.StartTime)
	}

	if task.EndTime != nil {
		sql.WriteString(" AND ctime<=?")
		args = append(args, task.EndTime)
	}

	v := Content{}

	count, err := db.Count(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	return &CountData{Total: int32(count)}, nil
}
