package wx

import (
	"encoding/base64"
	"fmt"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/http"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AppUp(app micro.IContext, task *AppUpTask) (*AppUpData, error) {

	if task.Type == nil {
		task.Type = "image"
	}

	b, err := base64.StdEncoding.DecodeString(task.Content)

	if err != nil {
		return nil, err
	}

	mp := dynamic.Get(app.GetConfig(), "mp")

	maxCount := 3
	forceUpdate := false

	var ret interface{} = nil

	for maxCount > 0 {

		err := func() error {

			token, err := MP_GetAccessToken(app, UserType_APP, task.Appid, forceUpdate)

			if err != nil {
				return err
			}

			ret, err = MP_SendWithType(app, mp,
				"POST",
				fmt.Sprintf("/cgi-bin/media/upload?access_token=%s&type=%s", token, task.Type),
				map[string]interface{}{"media": map[string]interface{}{"name": task.Name, "content": b}},
				http.OptionTypeMultipart)

			if err != nil {
				return err
			}

			return nil
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

		return nil, err
	}

	return &AppUpData{Type: dynamic.StringValue(dynamic.Get(ret, "type"), ""),
		MediaId:   dynamic.StringValue(dynamic.Get(ret, "media_id"), ""),
		CreatedAt: dynamic.IntValue(dynamic.Get(ret, "created_at"), 0)}, nil

}
