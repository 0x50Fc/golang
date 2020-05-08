package main

import (
	_ "github.com/go-sql-driver/mysql"
	S "github.com/hailongz/golang/features/addup/addup"
	"github.com/hailongz/golang/micro"
	"github.com/hailongz/golang/xs/def"
)

func In(config interface{}, s ...def.IService) error {

	srv := def.GetService(def.HTTP, s...).(def.IHTTPService)

	if srv != nil {

		app, err := micro.NewAppWithConfig(config, []micro.Service{&S.Service{}})

		if err != nil {
			return err
		}

		srv.HandleFunc("/"+app.GetName()+"/", micro.HandleFunc(app, nil), app)
	}

	return nil

}
