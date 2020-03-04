package wx

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) OpenMpTicketSet(app micro.IContext, task *OpenMpTicketSetTask) (*Open, error) {

	appid := dynamic.StringValue(dynamic.GetWithKeys(app.GetConfig(), []string{"open", "appid"}), "")

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Open{}

	err = db.Transaction(conn, func(conn db.Database) error {

		_, err := db.Get(conn, &v, prefix, " WHERE appid=? ", appid)

		if err != nil {
			return err
		}

		v.Appid = appid
		v.Ticket = task.Ticket

		if v.Id == 0 {
			_, err = db.Insert(conn, &v, prefix)
			if err != nil {
				return err
			}
		} else {
			_, err = db.UpdateWithKeys(conn, &v, prefix, map[string]bool{"ticket": true})
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	{
		redis, prefix, err := app.GetRedis("default")

		if err == nil {
			redis.Set(prefix+"open_ticket_"+appid, task.Ticket, time.Second*time.Duration(60*10)).Result()
		}
	}

	return nil, nil
}
