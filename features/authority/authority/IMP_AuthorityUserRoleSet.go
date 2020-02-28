package authority

import (
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AuthorityUserRoleSet(app micro.IContext, task *AuthorityUserRoleSetTask) (interface{}, error) {

	roleSet := map[int64]*Role{}

	if task.Role != "" {

		roleTask := RoleGetTask{}

		for _, s := range strings.Split(task.Role, ",") {
			roleTask.Name = s
			role, err := S.RoleGet(app, &roleTask)
			if err != nil {
				return nil, err
			}
			roleSet[role.Id] = role
		}
	}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Authority{}

	err = db.Transaction(conn, func(conn db.Database) error {

		items := []*Authority{}

		rs, err := db.Query(conn, &v, prefix, " WHERE uid=? AND resId=0", task.Uid)

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

			_, ok := roleSet[item.RoleId]

			if ok {
				delete(roleSet, item.RoleId)
			} else {
				_, err = db.Delete(conn, item, prefix)
				if err != nil {
					return err
				}
			}
		}

		for _, role := range roleSet {
			a := Authority{}
			a.RoleId = role.Id
			a.Uid = task.Uid
			app.Println(a)
			_, err = db.Insert(conn, &a, prefix)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return nil, nil

}
