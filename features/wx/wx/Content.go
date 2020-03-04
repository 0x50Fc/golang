package wx

import (
	"github.com/hailongz/golang/db"
)

type Content struct {
	db.Object
	Talker                 string `json:"talker" name:"talker" title:"微信原始数据库 的 talker"`
	Content                string `json:"content" name:"content" title:"微信原始数据库 的 content 消息内容" length:"512"`
	CreateTime             string `json:"createTime" name:"createtime" title:"微信原始数据库 的 createTime 消息时间(毫秒)"`
	Ctime                  int32  `json:"ctime" name:"ctime" title:"处理后的消息时间"`
	Etime                  int64  `json:"etime" name:"etime" title:"入库时间"`
	Type                   string `json:"type" name:"type" title:"微信原始数据库 的 type 判断消息类型"`
	IsSend                 string `json:"isSend" name:"issend" title:"微信原始数据库 的 isSend 判断是否是自己发送的消息"`
	MsgId                  string `json:"msgId" name:"msgid" title:"微信原始数据库的msgId 消息id自增" index:"ASC"`
	RobotUserAlias         string `json:"robotUserAlias" name:"robotuseralias" title:"机器人id" index:"ASC"`
	RobotUserName          string `json:"robotUserName" name:"robotusername" title:"机器人的微信用户名"`
	RobotUserNickName      string `json:"robotUserNickName" name:"robotusernickname" title:"机器人的微信昵称"`
	MsgTalkerUserName      string `json:"msgTalkerUserName" name:"msgtalkerusername" title:"消息发送者的微信用户名"`
	MsgTalkerUserAlias     string `json:"msgTalkerUserAlias" name:"msgtalkeruseralias" title:"消息发送者的微信id"`
	MsgTalkerUserNickName  string `json:"msgTalkerUserNickName" name:"msgtalkerusernickname" title:"消息发送者的微信昵称"`
	MsgTalkerUserReserved1 string `json:"msgTalkerUserReserved1" name:"msgtalkeruserreserved1" title:"消息发送者的微信头像大图"`
	MsgTalkerUserReserved2 string `json:"msgTalkerUserReserved2" name:"msgtalkeruserreserved2" title:"消息发送者的微信头像小图"`
	MsgGroupName           string `json:"msgGroupName" name:"msggroupname" title:"微信群的微信用户名"`
	MsgGroupNickName       string `json:"msgGroupNickName" name:"msggroupnickname" title:"微信群的微信昵称"`
	MsgContent             string `json:"msgContent" name:"msgcontent" title:"处理过的聊天内容"`
}

func (O *Content) GetName() string {
	return "content"
}

func (O *Content) GetTitle() string {
	return "群内容"
}
