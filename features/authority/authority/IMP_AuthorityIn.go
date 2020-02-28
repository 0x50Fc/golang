package authority

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AuthorityIn(app micro.IContext, task *AuthorityInTask) (interface{}, error) {

	resTask := ResGetTask{}

	resTask.Path = task.Path

	res, err := S.ResGet(app, &resTask)

	if err != nil {
		return nil, err
	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	v := Authority{}

	tbname := db.TableName(prefix, &v)
	sql := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE (uid=? AND resid=? AND roleid=0) OR (uid=? AND resid=0 AND roleid IN (SELECT b.roleid FROM %s as b WHERE b.uid=0 AND b.resid=?) )", tbname, tbname)

	rs, err := conn.Query(sql, task.Uid, res.Id, task.Uid, res.Id)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	if rs.Next() {
		n := int64(0)
		err = rs.Scan(&n)
		if err != nil {
			return nil, err
		}
		if n > 0 {
			return nil, nil
		}
	}

	return nil, micro.NewError(ERROR_NO_PERMISSION, "权限不足")
}
