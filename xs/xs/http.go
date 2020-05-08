package xs

import (
	"net/http"

	"github.com/hailongz/golang/xs/def"
)

type HttpService struct {
	Mux   *http.ServeMux
	items []def.IRecycle
}

func NewHttpService() *HttpService {
	return &HttpService{Mux: http.NewServeMux(), items: []def.IRecycle{}}
}

func (S *HttpService) Recycle() {
	for _, item := range S.items {
		item.Recycle()
	}
}

func (S *HttpService) Type() string {
	return def.HTTP
}

func (S *HttpService) HandleFunc(pattern string, handler func(resp http.ResponseWriter, req *http.Request), recycle def.IRecycle) {
	if recycle != nil {
		S.items = append(S.items, recycle)
	}
	S.Mux.HandleFunc(pattern, handler)
}

type HttpSrv struct {
	s *HttpService
}

func (sh *HttpSrv) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if sh.s != nil {
		sh.s.Mux.ServeHTTP(rw, req)
	} else {
		rw.WriteHeader(500)
		rw.Write([]byte{})
	}
}

func (sh *HttpSrv) SetService(s *HttpService) {
	if sh.s != s {
		if sh.s != nil {
			sh.s.Recycle()
		}
		sh.s = s
	}
}

func (sh *HttpSrv) IsEmpty() bool {
	return sh.s == nil
}
