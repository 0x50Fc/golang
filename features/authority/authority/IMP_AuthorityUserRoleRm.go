package authority

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AuthorityUserRoleRm(app micro.IContext, task *AuthorityUserRoleRmTask) (*Authority, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Authority{}

	err = db.Transaction(conn, func(conn db.Database) error {

		p, err := db.Get(conn, &v, prefix, " WHERE uid=? AND roleid=? AND resid=0", task.Uid, task.RoleId)

		if err != nil {
			return err
		}

		if p == nil {
			return micro.NewError(ERROR_NOT_FOUND, "未找到授权")
		}

		_, err = db.Delete(conn, &v, prefix)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &v, nil
}
