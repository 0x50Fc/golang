package adv

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Add(app micro.IContext, task *AddTask) (*Adv, error) {
	conn, prefix, err := app.GetDB("wd")
	if err != nil {
		return nil, err
	}
	adv := Adv{}
	adv.Title = dynamic.StringValue(task.Title, "")
	adv.Channel = dynamic.StringValue(task.Channel, "")
	adv.Position = int32(dynamic.IntValue(task.Position, 0))
	adv.Description = dynamic.StringValue(task.Description, "")
	adv.Pic = dynamic.StringValue(task.Pic, "")
	adv.Link = dynamic.StringValue(task.Link, "")
	adv.Starttime = dynamic.IntValue(task.Starttime, 0)
	adv.Endtime = dynamic.IntValue(task.Endtime, 0)
	adv.Linktype = int32(dynamic.IntValue(task.Linktype, 0))
	adv.Sort = task.Sort
	adv.Ctime = time.Now().Unix()

	_, err = db.Insert(conn, &adv, prefix)
	if err != nil {
		return nil, err
	}
	return &adv, nil
}
