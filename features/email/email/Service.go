package email

import (
	"git.sc.weibo.com/kk/microservice/golang/micro"
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
