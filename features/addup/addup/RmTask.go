package addup

type RmTask struct {
	Name	string	`json:"name" name:"name" title:"统计表名"`
	Region	interface{}	`json:"region,omitempty" name:"region" title:"分区 默认无分区"`
	Where	interface{}	`json:"where,omitempty" name:"where" title:"查询条件"`
	OrderBy	interface{}	`json:"orderBy,omitempty" name:"orderby" title:"排序"`
	Limit	interface{}	`json:"limit,omitempty" name:"limit" title:"限制条件"`
	Args	interface{}	`json:"args,omitempty" name:"args" title:"参数 JSON \n[\"\",123,245]"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "删除"
}

