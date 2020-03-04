package job

type SlaveQueryData struct {
	Items	[]*Slave	`json:"items,omitempty" name:"items" title:"主机"`
	Page	*Page	`json:"page,omitempty" name:"page" title:"分页"`
}

