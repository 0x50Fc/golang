package feed

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
	"git.sc.weibo.com/kk/microservice/id/client"
)

func (S *Service) OutboxSend(app micro.IContext, task *OutboxSendTask) (*Outbox, error) {

	conn, fix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	cli, err := micro.GetClient(app, "kk-id")

	if err != nil {
		return nil, err
	}

	id, err := client.API_Get(cli, &client.GetTask{})

	if err != nil {
		return nil, err
	}

	prefix := Prefix(app, fix, task.Uid)

	v := Outbox{}

	err = db.Transaction(conn, func(conn db.Database) error {

		err := func() error {

			rs, err := db.Query(conn, &v, prefix, " WHERE id=?", task.Id)

			if err != nil {
				return err
			}

			defer rs.Close()

			if rs.Next() {

				scaner := db.NewScaner(&v)

				err = scaner.Scan(rs)

				if err != nil {
					return err
				}

			} else {
				return micro.NewError(ERROR_NOT_FOUND, "未找到草稿")
			}

			return nil
		}()

		if err != nil {
			return err
		}

		if v.Status != OutboxStatus_None {
			return micro.NewError(ERROR_HAS_SENDED, "草稿已发布")
		}

		feed := Feed{}
		feed.Id = id
		feed.Uid = v.Uid
		feed.Body = v.Body
		feed.Options = v.Options
		feed.Ctime = time.Now().Unix()

		_, err = db.Insert(conn, &feed, Prefix(app, fix, id))

		if err != nil {
			return err
		}

		v.Mid = id
		v.Status = OutboxStatus_Sended
		v.Ctime = feed.Ctime

		_, err = db.UpdateWithKeys(conn, &v, prefix, map[string]bool{"status": true, "mid": true, "ctime": true})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
