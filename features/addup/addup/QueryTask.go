package addup

type QueryTask struct {
	Name	string	`json:"name" name:"name" title:"统计表名"`
	Region	interface{}	`json:"region,omitempty" name:"region" title:"分区 默认无分区"`
	Fields	interface{}	`json:"fields,omitempty" name:"fields" title:"查询字段 SQL , 默认 *"`
	Where	interface{}	`json:"where,omitempty" name:"where" title:"查询条件"`
	OrderBy	interface{}	`json:"orderBy,omitempty" name:"orderby" title:"排序"`
	GroupBy	interface{}	`json:"groupBy,omitempty" name:"groupby" title:"分组"`
	Having	interface{}	`json:"having,omitempty" name:"having" title:"分组筛选条件"`
	Limit	interface{}	`json:"limit,omitempty" name:"limit" title:"限制条件"`
	Args	interface{}	`json:"args,omitempty" name:"args" title:"参数 JSON \n[\"\",123,245]"`
	CacheKey	interface{}	`json:"cacheKey,omitempty" name:"cachekey" title:"缓存"`
}

func (T *QueryTask) GetName() string {
	return "query.json"
}

func (T *QueryTask) GetTitle() string {
	return "查询"
}

