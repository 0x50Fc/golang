package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/hailongz/golang/features/oss/ali"
	_ "github.com/hailongz/golang/features/oss/minio"
	S "github.com/hailongz/golang/features/oss/oss"
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
