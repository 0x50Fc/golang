package wx

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
	"time"
)

func (S *Service) ContentRecv(app micro.IContext, task *ContentRecvTask) (*Content, error) {
	conn, prefix, err := app.GetDB("wd")
	if err != nil {
		return nil, err
	}

	content := Content{}
	content.Talker = task.Talker
	content.Content = task.Content
	content.CreateTime = task.CreateTime
	content.Ctime = task.Ctime
	content.Etime = time.Now().Unix()
	content.IsSend = task.IsSend
	content.MsgContent = task.MsgContent
	content.MsgGroupName = task.MsgGroupName
	content.MsgGroupNickName = task.MsgGroupNickName
	content.MsgId = task.MsgId
	content.MsgTalkerUserAlias = task.MsgTalkerUserAlias
	content.MsgTalkerUserName = task.MsgTalkerUserName
	content.MsgTalkerUserNickName = task.MsgTalkerUserNickName
	content.MsgTalkerUserReserved1 = task.MsgTalkerUserReserved1
	content.MsgTalkerUserReserved2 = task.MsgTalkerUserReserved2
	content.RobotUserAlias = task.RobotUserAlias
	content.RobotUserName = task.RobotUserName
	content.RobotUserNickName = task.RobotUserNickName
	content.Type = task.Type
	_, err = db.Insert(conn, &content, prefix)
	if err != nil {
		return nil, err
	}
	return &content, nil
}
