package xs

import (
	"net/http"
	"strings"

	"github.com/hailongz/golang/xs/def"
)

type ServiceItem struct {
	prefix  string
	recycle def.IRecycle
	handler func(resp http.ResponseWriter, req *http.Request)
}

type Service struct {
	items []*ServiceItem
}

func NewService() *Service {
	return &Service{items: []*ServiceItem{}}
}

func (S *Service) HandleFunc(prefix string, recycle def.IRecycle, handler func(resp http.ResponseWriter, req *http.Request)) {
	S.items = append(S.items, &ServiceItem{prefix: prefix, recycle: recycle, handler: handler})
}

func (S *Service) Handle(resp http.ResponseWriter, req *http.Request) bool {

	for _, item := range S.items {
		if !strings.HasPrefix(req.URL.Path, item.prefix) {
			continue
		}
		item.handler(resp, req)
		return true
	}

	return false
}

func (S *Service) Recycle() {
	for _, item := range S.items {
		if item.recycle != nil {
			item.recycle.Recycle()
		}
	}
}
