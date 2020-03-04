package wx

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) CheckText(app micro.IContext, task *CheckTextTask) (interface{}, error) {

	mp := dynamic.Get(app.GetConfig(), "mp")

	if task.Content == "" {
		return nil, micro.NewError(ERROR_BODY, "消息内容格式错误")
	}

	data := map[string]string{}
	data["content"] = task.Content

	//获取token
	maxCount := 3
	forceUpdate := false

	for maxCount > 0 {

		ret, err := func() (interface{}, error) {
			token, err := MP_GetAccessToken(app, UserType_APP, task.Appid, forceUpdate)

			if err != nil {
				return nil, err
			}

			return MP_Send(app, mp, "POST", "/wxa/msg_sec_check?access_token="+token, data)

		}()

		if err == nil {
			break
		}

		e, ok := err.(*micro.Error)

		if ok {
			if e.Errno == 40001 {
				maxCount = maxCount - 1
				forceUpdate = true
				continue
			}
		}

		return ret, err
	}

	return nil, nil
}
