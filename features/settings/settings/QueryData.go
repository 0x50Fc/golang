package settings

type QueryData struct {
	Items	[]*Setting	`json:"items,omitempty" name:"items" title:"配置数据"`
	Page	*Page	`json:"page,omitempty" name:"page" title:"分页"`
}

