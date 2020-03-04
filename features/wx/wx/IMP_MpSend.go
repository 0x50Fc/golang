package wx

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) MpSend(app micro.IContext, task *MpSendTask) (interface{}, error) {

	mp := dynamic.Get(app.GetConfig(), "mp")

	var data interface{} = nil

	err := json.Unmarshal([]byte(task.Body), &data)

	if err != nil {
		return nil, micro.NewError(ERROR_BODY, "消息内容格式错误")
	}

	maxCount := 3
	forceUpdate := false

	for maxCount > 0 {

		err := func() error {

			token, err := MP_GetAccessToken(app, UserType_MP, task.Appid, forceUpdate)

			if err != nil {
				return err
			}

			body := map[string]interface{}{}
			body["touser"] = task.Openid

			if task.KfAccount != nil {
				body["customservice"] = map[string]interface{}{"kf_account": task.KfAccount}
			}

			if task.Type == MessageType_Image {
				body["msgtype"] = "image"
				body["image"] = data
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token, body)
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_Voice {
				body["msgtype"] = "voice"
				body["voice"] = data
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token, body)
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_Video {
				body["msgtype"] = "video"
				body["video"] = data
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token, body)
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_Music {
				body["msgtype"] = "music"
				body["music"] = data
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token, body)
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_News {
				body["msgtype"] = "news"
				body["news"] = data
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token, body)
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_Mpnews {
				body["msgtype"] = "mpnews"
				body["mpnews"] = data
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token, body)
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_Msgmenu {
				body["msgtype"] = "msgmenu"
				body["msgmenu"] = data
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token, body)
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_Wxcard {
				body["msgtype"] = "wxcard"
				body["wxcard"] = data
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token, body)
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_Miniprogrampage {
				body["msgtype"] = "miniprogrampage"
				body["miniprogrampage"] = data
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token, body)
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_Template {
				dynamic.Each(data, func(key interface{}, value interface{}) bool {
					body[dynamic.StringValue(key, "")] = value
					return true
				})
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/template/send?access_token="+token, body)
				if err != nil {
					return err
				}
			} else {
				body["msgtype"] = "text"
				body["text"] = data
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token, body)
				if err != nil {
					return err
				}
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

	return nil, nil
}
