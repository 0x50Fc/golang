package wx

type ContentRmTask struct {
	Id      int64       `json:"id" name:"id" title:"内容ID"`
	Uid     interface{} `json:"uid,omitempty" name:"uid" title:"用户ID"`
	Groupid interface{} `json:"groupid,omitempty" name:"groupid" title:"群ID"`
}

func (T *ContentRmTask) GetName() string {
	return "content/rm.json"
}

func (T *ContentRmTask) GetTitle() string {
	return "删除评论"
}
