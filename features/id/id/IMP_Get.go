package id

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/iid"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Get(app micro.IContext, task *GetTask) (int64, error) {

	v, err := app.GetSharedObject("iid", func() (micro.SharedObject, error) {
		return iid.NewIID(dynamic.IntValue(dynamic.Get(app.GetConfig(), "aid"), 0),
			dynamic.IntValue(dynamic.Get(app.GetConfig(), "nid"), 0)), nil
	})

	if err != nil {
		return 0, err
	}

	iid := v.(*iid.IID)

	return -iid.NewID(), nil
}
