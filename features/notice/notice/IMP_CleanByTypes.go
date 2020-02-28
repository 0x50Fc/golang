package notice

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) CleanByTypes(app micro.IContext, task *CleanByTypesTask) (interface{}, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	tableCount := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"table", "count"}), 128)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	v := Notice{}

	sql.WriteString(fmt.Sprintf("DELETE FROM %s{{index}}_%s WHERE fid=? AND iid = ?", prefix, v.GetName()))
	args = append(args, task.Fid, task.Iid)

	vs := strings.Split(dynamic.StringValue(task.Type, ""), ",")

	sql.WriteString(" AND type IN (")

	for j, vv := range vs {
		if j != 0 {
			sql.WriteString(",")
		}
		sql.WriteString("?")
		args = append(args, vv)
	}

	sql.WriteString("); ")

	app.Println(sql.String())

	err = db.Transaction(conn, func(conn db.Database) error {

		for i := int64(1); i <= tableCount; i++ {

			_, err := conn.Exec(strings.Replace(sql.String(), "{{index}}", fmt.Sprintf("%d", i), 1), args...)

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}
