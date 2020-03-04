package wx

import (
	"fmt"
	"net/url"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) OpenMpAuthorize(app micro.IContext, task *OpenMpAuthorizeTask) (*OpenMPAuthorizeData, error) {

	mp := dynamic.Get(app.GetConfig(), "mp")

	appid := dynamic.StringValue(dynamic.GetWithKeys(app.GetConfig(), []string{"open", "appid"}), "")

	token, err := MP_Open_GetAccessToken(app, appid)

	if err != nil {
		return nil, err
	}

	ret, err := MP_Send(app, mp, "POST", "/cgi-bin/component/api_create_preauthcode?component_access_token="+token, map[string]interface{}{"component_appid": appid})

	if err != nil {
		return nil, err
	}

	pre_auth_code := dynamic.StringValue(dynamic.Get(ret, "pre_auth_code"), "")
	authType := dynamic.IntValue(task.AuthType, OpenMPAuthType_MP_APP)
	openType := dynamic.IntValue(task.OpenType, OpenMPOpenType_WX)

	if openType == OpenMPOpenType_Web {
		u := fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s",
			appid,
			pre_auth_code,
			url.QueryEscape(dynamic.StringValue(task.RedirectUri, "")))
		if task.AuthType != nil {
			u = fmt.Sprintf("%s&auth_type=%d", u, authType)
		}
		if task.Appid != nil {
			u = fmt.Sprintf("%s&biz_appid=%s", u, task.Appid)
		}
		return &OpenMPAuthorizeData{Url: u}, nil
	} else {
		u := fmt.Sprintf("https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s",
			appid,
			pre_auth_code,
			url.QueryEscape(dynamic.StringValue(task.RedirectUri, "")))
		if task.AuthType != nil {
			u = fmt.Sprintf("%s&auth_type=%d", u, authType)
		}
		if task.Appid != nil {
			u = fmt.Sprintf("%s&biz_appid=%s", u, task.Appid)
		}
		u = u + "#wechat_redirect"
		return &OpenMPAuthorizeData{Url: u}, nil
	}

}
