package main

import (
	_ "github.com/go-sql-driver/mysql"
	S "github.com/hailongz/golang/features/adv/adv"
	"github.com/hailongz/golang/micro"
	"github.com/hailongz/golang/xs/def"
)

func In(config interface{}, s def.IService) error {

	app, err := micro.NewAppWithConfig(config, []micro.Service{&S.Service{}})

	if err != nil {
		return err
	}

	s.HandleFunc("/"+app.GetName()+"/", app, micro.HandleFunc(app, nil))

	return nil

}
