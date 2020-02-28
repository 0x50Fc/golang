package notice

type CleanByTypesTask struct {
	Type string `json:"type" name:"type" title:"类型, 多个逗号分割"`
	Fid  int64  `json:"fid" name:"fid" title:"消息来源ID , 多个逗号分割"`
	Iid  int64  `json:"iid" name:"iid" title:"消息来源项ID , 多个逗号分割"`
}

func (T *CleanByTypesTask) GetName() string {
	return "cleanByTypes.json"
}

func (T *CleanByTypesTask) GetTitle() string {
	return "清理消息"
}
