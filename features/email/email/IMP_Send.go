package email

import (
	"github.com/hailongz/golang//dynamic"
	"github.com/hailongz/golang//micro"
	"gopkg.in/gomail.v2"
)

func (S *Service) Send(app micro.IContext, task *SendTask) (interface{}, error) {

	config := dynamic.Get(app.GetConfig(), "email")
	from := dynamic.StringValue(dynamic.Get(config, "from"), "")
	addr := dynamic.StringValue(dynamic.Get(config, "addr"), "exmail.qq.com")
	port := int(dynamic.IntValue(dynamic.Get(config, "port"), 587))
	user := dynamic.StringValue(dynamic.Get(config, "user"), from)
	password := dynamic.StringValue(dynamic.Get(config, "password"), "")

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", task.To)
	m.SetHeader("Subject", task.Subject)
	m.SetBody(task.ContentType, task.Body)

	d := gomail.NewDialer(addr, port, user, password)

	if err := d.DialAndSend(m); err != nil {
		return nil, err
	}

	return nil, nil
}
