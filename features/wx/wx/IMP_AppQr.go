package wx

import (
	"encoding/base64"
	"fmt"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/http"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AppQr(app micro.IContext, task *AppQrTask) (*WXAppQRData, error) {

	mp := dynamic.Get(app.GetConfig(), "mp")

	maxCount := 3
	forceUpdate := false

	var result interface{} = nil

	data := map[string]interface{}{}
	data["scene"] = task.Scene

	if task.Page != nil {
		data["page"] = task.Page
	}

	if task.Width != nil {
		data["width"] = dynamic.IntValue(task.Width, 0)
	}

	data["is_hyaline"] = true

	for maxCount > 0 {

		err := func() error {

			token, err := MP_GetAccessToken(app, UserType_APP, task.Appid, forceUpdate)

			if err != nil {
				return err
			}

			dynamic.Each(dynamic.Get(mp, "baseURL"), func(_ interface{}, baseURL interface{}) bool {

				options := http.Options{}
				options.Method = "POST"
				options.Url = fmt.Sprintf("%s/wxa/getwxacodeunlimit?access_token=%s", baseURL, token)
				options.Data = data
				options.ResponseType = http.OptionResponseTypeByte
				options.Type = http.OptionTypeJson

				app.Println("[MP_Send]", options.Url, options.Data)

				result, err = http.Send(&options)

				if err != nil || result == nil {
					return true
				}

				var rs interface{} = nil

				err = json.Unmarshal(result.([]byte), &rs)

				if err != nil {
					err = nil
					return false
				}

				app.Println("[MP_Send]", rs, err)

				errcode := dynamic.IntValue(dynamic.Get(rs, "errcode"), 0)

				if errcode == -1 {
					return false
				}

				if errcode != 0 {
					err = micro.NewError(int(errcode), dynamic.StringValue(dynamic.Get(rs, "errmsg"), "未知微信服务错误"))
					return false
				}

				return false
			})

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

	b, ok := result.([]byte)

	if !ok {
		return nil, micro.NewError(ERROR_QR, "无法生成小程序二维码")
	}

	return &WXAppQRData{Type: "base64", Content: base64.StdEncoding.EncodeToString(b)}, nil
}
