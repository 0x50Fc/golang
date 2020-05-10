package def

import "net/http"

type IRecycle interface {
	Recycle()
}

type IService interface {
	HandleFunc(pattern string, recycle IRecycle, handler func(resp http.ResponseWriter, req *http.Request))
}

type In = func(config interface{}, s IService) error
