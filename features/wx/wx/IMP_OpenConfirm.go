package wx

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) OpenConfirm(app micro.IContext, task *OpenConfirmTask) (*User, error) {

	mp := dynamic.Get(app.GetConfig(), "mp")

	appid := dynamic.StringValue(dynamic.GetWithKeys(app.GetConfig(), []string{"open", "appid"}), "")

	token, err := MP_Open_GetAccessToken(app, appid)

	if err != nil {
		return nil, err
	}

	data, err := MP_Send(app, mp, "GET", "/sns/oauth2/component/access_token",
		map[string]interface{}{
			"appid":                  task.Appid,
			"component_appid":        appid,
			"component_access_token": token,
			"code":       task.Code,
			"grant_type": "authorization_code"},
	)

	if err != nil {
		return nil, err
	}

	openid := dynamic.StringValue(dynamic.Get(data, "openid"), "")
	access_token := dynamic.StringValue(dynamic.Get(data, "access_token"), "")
	refresh_token := dynamic.StringValue(dynamic.Get(data, "refresh_token"), "")
	expires_in := dynamic.IntValue(dynamic.Get(data, "expires_in"), 0)
	unionid := dynamic.StringValue(dynamic.Get(data, "unionid"), "")

	v := User{}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()

	err = db.Transaction(conn, func(conn db.Database) error {

		_, err := db.Get(conn, &v, prefix, " WHERE type=? AND appid=? AND openid=?", UserType_MP, task.Appid, openid)

		if err != nil {
			return err
		}

		v.Type = UserType_MP
		v.Appid = task.Appid
		v.AccessToken = access_token
		v.RefreshToken = refresh_token
		v.Etime = now + expires_in - 30
		v.Openid = openid
		v.Unionid = unionid
		v.Ctime = now

		if v.Id == 0 {
			_, err = db.Insert(conn, &v, prefix)
			if err != nil {
				return err
			}
		} else {
			_, err = db.UpdateWithKeys(conn, &v, prefix, map[string]bool{"access_token": true, "refresh_token": true, "etime": true, "unionid": true})
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	err = MP_UpdateUser(app, &v)

	if err != nil {
		db.Update(conn, &v, prefix)
	}

	return &v, nil
}
