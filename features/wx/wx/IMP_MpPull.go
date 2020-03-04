package wx

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) MpPull(app micro.IContext, task *MpPullTask) (*WXMPPullData, error) {

	mp := dynamic.Get(app.GetConfig(), "mp")

	token, err := MP_GetAccessToken(app, UserType_MP, task.Appid, false)

	if err != nil {
		return nil, err
	}

	var count int32 = 0
	var total int32 = 0

	next_openid := ""

	for {

		data, err := MP_Send(app, mp, "GET", "/cgi-bin/user/get", map[string]interface{}{"access_token": token, "next_openid": next_openid})

		if err != nil {
			return nil, err
		}

		n := int32(dynamic.IntValue(dynamic.Get(data, "count"), 0))

		if n == 0 {
			break
		}

		total = int32(dynamic.IntValue(dynamic.Get(data, "total"), int64(total)))

		next_openid = dynamic.StringValue(dynamic.Get(data, "next_openid"), next_openid)

		dynamic.Each(dynamic.Get(data, ""), func(_ interface{}, openid interface{}) bool {

			get := GetTask{}

			get.Appid = task.Appid
			get.Type = UserType_MP
			get.Openid = dynamic.StringValue(openid, "")

			S.Get(app, &get)

			return true
		})

		count = count + n

		if count >= total {
			break
		}
	}

	return &WXMPPullData{Count: count}, nil
}
