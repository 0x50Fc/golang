package member

type QueryData struct {
	Items	[]*Member	`json:"items,omitempty" name:"items" title:"成员" jsonType:"true"`
	Page	*QueryDataPage	`json:"page,omitempty" name:"page" title:"分页" jsonType:"true"`
}

