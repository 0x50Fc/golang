package wx

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/http"
	"github.com/hailongz/golang/micro"
	"path"
)

func (S *Service) CheckImage(app micro.IContext, task *CheckImageTask) (interface{}, error) {

	mp := dynamic.Get(app.GetConfig(), "mp")

	if task.Url == "" {
		return nil, micro.NewError(ERROR_BODY, "图片格式错误")
	}

	data := map[string]interface{}{}

	opt := http.Options{}
	opt.Method = "GET"
	opt.Url = task.Url
	opt.ResponseType = http.OptionResponseTypeByte

	media, err := http.Send(&opt)

	if err != nil {
		return nil, err
	}

	fileName := path.Base(task.Url)
	data["media"] = map[string]interface{}{"name": fileName, "content": media.([]byte)}

	//获取token
	//获取token
	maxCount := 3
	forceUpdate := false

	for maxCount > 0 {

		ret, err := func() (interface{}, error) {
			token, err := MP_GetAccessToken(app, UserType_APP, task.Appid, forceUpdate)

			if err != nil {
				return nil, err
			}

			return MP_SendWithType(app, mp, "POST", "/wxa/img_sec_check?access_token="+token, data, http.OptionTypeMultipart)

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
