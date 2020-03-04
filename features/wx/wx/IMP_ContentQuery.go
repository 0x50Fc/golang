package wx

import (
	"bytes"
	"fmt"
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) ContentQuery(app micro.IContext, task *ContentQueryTask) (*ContentList, error) {

	p := int32(dynamic.IntValue(task.P, 1))
	n := int32(dynamic.IntValue(task.N, 20))

	if n < 1 {
		n = 20
	}

	if p < 1 {
		p = 1
	}

	data := ContentList{}

	if task.P != nil {

		countTask := ContentCountTask{}

		countTask.MsgId = task.MsgId
		countTask.CreateTime = task.CreateTime
		countTask.MsgTalkerUserAlias = task.MsgTalkerUserAlias
		countTask.MsgTalkerUserNickName = task.MsgTalkerUserNickName
		countTask.Type = task.Type
		countTask.MsgGroupName = task.MsgGroupName
		countTask.Q = task.Q
		countTask.StartTime = task.StartTime
		countTask.EndTime = task.EndTime

		countData, err := S.ContentCount(app, &countTask)

		if err != nil {
			return nil, err
		}

		data.Page = &Page{
			Total: countData.Total,
			P:     p,
			N:     n,
			Count: countData.Total / n,
		}

		if countData.Total%n != 0 {
			data.Page.Count++
		}
	}

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

	sql.WriteString(fmt.Sprintf(" ORDER BY id DESC LIMIT %d,%d", (p-1)*n, n))

	v := Content{}

	data.Items = []*Content{}

	rs, err := db.Query(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	scaner := db.NewScaner(&v)

	for rs.Next() {

		err = scaner.Scan(rs)

		if err != nil {
			return nil, err
		}

		item := Content{}
		item = v

		data.Items = append(data.Items, &item)
	}

	return &data, nil
}
