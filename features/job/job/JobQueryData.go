package job

type JobQueryData struct {
	Items	[]*Job	`json:"items,omitempty" name:"items" title:"工作"`
	Page	*Page	`json:"page,omitempty" name:"page" title:"分页"`
}

