package main

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
