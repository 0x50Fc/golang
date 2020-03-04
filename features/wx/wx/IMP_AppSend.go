package wx

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AppSend(app micro.IContext, task *AppSendTask) (interface{}, error) {

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

			token, err := MP_GetAccessToken(app, UserType_APP, task.Appid, forceUpdate)

			if err != nil {
				return err
			}

			if task.Type == MessageType_Image {
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token,
					map[string]interface{}{"touser": task.Openid, "msgtype": "image", "image": data})
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_Voice {
				return micro.NewError(ERROR_MESSAGE_TYPE, "不支持的消息类型")
			} else if task.Type == MessageType_Video {
				return micro.NewError(ERROR_MESSAGE_TYPE, "不支持的消息类型")
			} else if task.Type == MessageType_Music {
				return micro.NewError(ERROR_MESSAGE_TYPE, "不支持的消息类型")
			} else if task.Type == MessageType_News {
				return micro.NewError(ERROR_MESSAGE_TYPE, "不支持的消息类型")
			} else if task.Type == MessageType_Mpnews {
				return micro.NewError(ERROR_MESSAGE_TYPE, "不支持的消息类型")
			} else if task.Type == MessageType_Msgmenu {
				return micro.NewError(ERROR_MESSAGE_TYPE, "不支持的消息类型")
			} else if task.Type == MessageType_Wxcard {
				return micro.NewError(ERROR_MESSAGE_TYPE, "不支持的消息类型")
			} else if task.Type == MessageType_Miniprogrampage {
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token,
					map[string]interface{}{"touser": task.Openid, "msgtype": "miniprogrampage", "miniprogrampage": data})
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_Link {
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token,
					map[string]interface{}{"touser": task.Openid, "msgtype": "link", "link": data})
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_Template {
				opt := map[string]interface{}{"touser": task.Openid}
				dynamic.Each(data, func(key interface{}, value interface{}) bool {
					opt[dynamic.StringValue(key, "")] = value
					return true
				})
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/wxopen/template/send?access_token="+token, opt)
				if err != nil {
					return err
				}
			} else if task.Type == MessageType_Subscribe_Template {
				opt := map[string]interface{}{"touser": task.Openid}
				dynamic.Each(data, func(key interface{}, value interface{}) bool {
					opt[dynamic.StringValue(key, "")] = value
					return true
				})
				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/subscribe/send?access_token="+token, opt)
				if err != nil {
					return err
				}
			} else {

				_, err := MP_Send(app, mp, "POST", "/cgi-bin/message/custom/send?access_token="+token,
					map[string]interface{}{"touser": task.Openid, "msgtype": "text", "text": data})
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
