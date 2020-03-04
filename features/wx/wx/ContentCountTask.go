package wx

type ContentCountTask struct {
	Id                    interface{} `json:"id,omitempty" name:"id" title:"ID"`
	MsgId                 string      `json:"msgId" name:"msgid" title:"微信原始数据库的msgId 消息id自增" index:"ASC"`
	CreateTime            string      `json:"createTime" name:"createtime" title:"微信原始数据库 的 createTime 消息时间(毫秒)"`
	MsgTalkerUserAlias    string      `json:"msgTalkerUserAlias" name:"msgtalkeruseralias" title:"消息发送者的微信id"`
	MsgTalkerUserNickName string      `json:"msgTalkerUserNickName" name:"msgtalkerusernickname" title:"消息发送者的微信昵称"`
	Type                  string      `json:"type" name:"type" title:"微信原始数据库 的 type 判断消息类型"`
	MsgGroupName          string      `json:"msgGroupName" name:"msggroupname" title:"微信群的微信用户名"`
	StartTime             interface{} `json:"startTime,omitempty" name:"starttime" title:"开始时间"`
	EndTime               interface{} `json:"endTime,omitempty" name:"endtime" title:"结束时间"`
	Q                     interface{} `json:"q,omitempty" name:"q" title:"模糊匹配关键字"`
	P                     interface{} `json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N                     interface{} `json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *ContentCountTask) GetName() string {
	return "content/count.json"
}

func (T *ContentCountTask) GetTitle() string {
	return "查询"
}
