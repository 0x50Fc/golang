package wx

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AppLogin(app micro.IContext, task *AppLoginTask) (*AppLoginData, error) {

	mp := dynamic.Get(app.GetConfig(), "mp")

	secret, err := MP_GetSecret(mp, task.Appid)

	if err != nil {
		return nil, err
	}

	ret, err := MP_Send(app, mp, "GET", "/sns/jscode2session", map[string]interface{}{"appid": task.Appid, "secret": secret, "grant_type": "authorization_code", "js_code": task.Code})

	if err != nil {
		return nil, err
	}

	openid := dynamic.StringValue(dynamic.Get(ret, "openid"), "")
	unionid := dynamic.StringValue(dynamic.Get(ret, "unionid"), "")
	session_key := dynamic.StringValue(dynamic.Get(ret, "session_key"), "")

	v := User{}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	_, err = db.Get(conn, &v, prefix, " WHERE type=? AND appid=? AND openid=?", UserType_APP, task.Appid, openid)

	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()

	v.Type = UserType_APP
	v.Appid = task.Appid
	v.Openid = openid

	if unionid != "" {
		v.Unionid = unionid
	}

	v.SessionKey = session_key
	v.Mtime = now

	if task.EncryptedData != nil && task.Iv != nil {

		err = func() error {

			ret, err := MP_AppDecrypt(app, session_key, dynamic.StringValue(task.EncryptedData, ""), dynamic.StringValue(task.Iv, ""))

			if err != nil {
				return err
			}

			var info interface{} = nil

			err = json.Unmarshal(ret, &info)

			if err != nil {
				app.Println("[json.Unmarshal]", err, string(ret))
				return err
			}

			app.Println("[WXUSER]", string(ret))

			if dynamic.StringValue(dynamic.Get(info, "openId"), "") == openid &&
				dynamic.StringValue(dynamic.GetWithKeys(info, []string{"watermark", "appid"}), "") == task.Appid {
				v.Nick = dynamic.StringValue(dynamic.Get(info, "nickName"), v.Nick)
				v.Logo = dynamic.StringValue(dynamic.Get(info, "avatarUrl"), v.Logo)
				v.Gender = int32(dynamic.IntValue(dynamic.Get(info, "gender"), int64(v.Gender)))
				v.Country = dynamic.StringValue(dynamic.Get(info, "country"), v.Country)
				v.Province = dynamic.StringValue(dynamic.Get(info, "province"), v.Province)
				v.City = dynamic.StringValue(dynamic.Get(info, "city"), v.City)
				v.Lang = dynamic.StringValue(dynamic.Get(info, "language"), v.Lang)
				v.Unionid = dynamic.StringValue(dynamic.Get(info, "unionId"), v.Unionid)
			}

			return nil
		}()

		if err != nil {
			return nil, err
		}

	}

	if v.Id == 0 {
		v.Ctime = now
		_, err = db.Insert(conn, &v, prefix)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = db.Update(conn, &v, prefix)
		if err != nil {
			return nil, err
		}
	}

	return &AppLoginData{SessionKey: session_key, User: &v}, nil
}
