package wx

import (
	"fmt"
	"net/url"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) MpAuthorize(app micro.IContext, task *MpAuthorizeTask) (*MPAuthorizeData, error) {

	u := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect",
		task.Appid, url.QueryEscape(task.RedirectUri), url.QueryEscape(dynamic.StringValue(task.Scope, Scope_BASE)), url.QueryEscape(dynamic.StringValue(task.State, "")))

	return &MPAuthorizeData{Url: u}, nil
}
