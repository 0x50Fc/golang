package micro

import (
	"errors"
	"fmt"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/http"
)

type Client interface {
	Recycle()
	Send(method string, name string, data interface{}) (interface{}, error)
}

type HTTPClient struct {
	baseURL string
}

func NewHttpClient(baseURL string) *HTTPClient {
	return &HTTPClient{baseURL: baseURL}
}

func (C *HTTPClient) Recycle() {

}

func (C *HTTPClient) Send(method string, name string, data interface{}) (interface{}, error) {

	opt := http.Options{}
	opt.Url = C.baseURL + name
	opt.Method = method
	opt.Data = data
	opt.Type = http.OptionTypeUrlencode
	opt.ResponseType = http.OptionResponseTypeJson

	data, err := http.Send(&opt)

	if err != nil {
		return nil, err
	}

	v := dynamic.Get(data, "errno")

	if v == nil {
		return nil, errors.New("服务接口数据错误 " + opt.Url)
	}

	errno := int(dynamic.IntValue(v, 0))

	if errno == 200 {
		return dynamic.Get(data, "data"), nil
	}

	return nil, NewError(errno, dynamic.StringValue(dynamic.Get(data, "errmsg"), "未知错误"))
}

func GetClient(app IContext, name string) (Client, error) {

	cfg := dynamic.Get(app.GetConfig(), "client")

	v, err := app.GetSharedObject(fmt.Sprintf("__client_%s", name), func() (SharedObject, error) {

		baseURL := ""

		v := dynamic.Get(cfg, name)

		if v == nil {
			baseURL = dynamic.StringValue(dynamic.Get(cfg, "baseURL"), "") + name + "/"
		} else {
			baseURL = dynamic.StringValue(v, "")
		}

		if baseURL == "" {
			return nil, errors.New("[Client] 未找到可用配置")
		}

		return NewHttpClient(baseURL), nil
	})

	if err != nil {
		return nil, err
	}

	return v.(Client), nil
}
