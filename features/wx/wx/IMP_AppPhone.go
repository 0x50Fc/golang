package wx

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AppPhone(app micro.IContext, task *AppPhoneTask) (*AppPhoneData, error) {

	var v *User = nil
	var err error = nil

	sessionKey := ""

	if task.SessionKey != nil {
		sessionKey = dynamic.StringValue(task.SessionKey, "")
	} else if task.Openid != nil {

		t := GetTask{}

		t.Appid = task.Appid
		t.Type = UserType_APP
		t.Openid = dynamic.StringValue(task.Openid, "")

		v, err = S.Get(app, &t)

		if err != nil {
			return nil, err
		}

		if v.SessionKey == "" {
			return nil, micro.NewError(ERROR_NOT_FOUND, "未找到 Session Key")
		}

		sessionKey = v.SessionKey

	}

	data, err := MP_AppDecrypt(app, sessionKey, task.EncryptedData, task.Iv)

	if err != nil {
		return nil, err
	}

	var info interface{} = nil

	err = json.Unmarshal(data, &info)

	if err != nil {
		app.Println("[json.Unmarshal]", err, string(data))
		return nil, err
	}

	if dynamic.StringValue(dynamic.GetWithKeys(info, []string{"watermark", "appid"}), "") == task.Appid {
		return &AppPhoneData{
			User:    v,
			Phone:   dynamic.StringValue(dynamic.Get(info, "purePhoneNumber"), ""),
			Country: dynamic.StringValue(dynamic.Get(info, "countryCode"), ""),
		}, nil
	}

	return nil, micro.NewError(ERROR_NOT_FOUND, "错误的编码数据")
}
