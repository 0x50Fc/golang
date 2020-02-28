package addup

import (
	"github.com/hailongz/golang/micro"
)

func init() {
	micro.AddDefaultService(&Service{})
}

type Service struct {
}

func (S *Service) GetName() string {
	return ""
}

func (S *Service) GetTitle() string {
	return "默认服务"
}
