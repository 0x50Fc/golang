package def

import "net/http"

const (
	HTTP = "http"
)

type IRecycle interface {
	Recycle()
}

type IService interface {
	IRecycle
	Type() string
}

type IHTTPService interface {
	IService
	HandleFunc(pattern string, handler func(resp http.ResponseWriter, req *http.Request), recycle IRecycle)
}

type In = func(config interface{}, s ...IService) error

func GetService(stype string, s ...IService) IService {
	if s != nil {
		for _, v := range s {
			if v.Type() == stype {
				return v
			}
		}
	}
	return nil
}
