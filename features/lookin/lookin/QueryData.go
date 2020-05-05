package lookin

type QueryData struct {
	Items	[]*Lookin	`json:"items,omitempty" name:"items" title:"在看" jsonType:"true"`
	Page	*QueryDataPage	`json:"page,omitempty" name:"page" title:"分页" jsonType:"true"`
}

