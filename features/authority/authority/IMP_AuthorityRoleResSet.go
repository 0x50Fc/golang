package authority

import (
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AuthorityRoleResSet(app micro.IContext, task *AuthorityRoleResSetTask) (interface{}, error) {

	resSet := map[int64]*Res{}

	if task.Res != "" {

		resTask := ResGetTask{}

		for _, s := range strings.Split(task.Res, ",") {
			resTask.Path = s
			res, err := S.ResGet(app, &resTask)
			if err != nil {
				return nil, err
			}
			resSet[res.Id] = res
		}
	}

	roleTask := RoleGetTask{}
	roleTask.Name = task.Role

	role, err := S.RoleGet(app, &roleTask)

	if err != nil {
		return nil, err
	}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Authority{}

	err = db.Transaction(conn, func(conn db.Database) error {

		items := []*Authority{}

		rs, err := db.Query(conn, &v, prefix, " WHERE uid=0 AND roleid=?", role.Id)

		if err != nil {
			return err
		}

		scaner := db.NewScaner(&v)

		for rs.Next() {

			err = scaner.Scan(rs)

			if err != nil {
				rs.Close()
				return err
			}

			item := Authority{}
			item = v

			items = append(items, &item)
		}

		rs.Close()

		for _, item := range items {

			_, ok := resSet[item.ResId]

			if ok {
				delete(resSet, item.ResId)
			} else {
				_, err = db.Delete(conn, item, prefix)
				if err != nil {
					return err
				}
			}
		}

		for _, res := range resSet {
			a := Authority{}
			a.RoleId = role.Id
			a.ResId = res.Id
			_, err = db.Insert(conn, &a, prefix)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return nil, nil
}
