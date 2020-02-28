package top

type QueryData struct {
	Items	[]*Top	`json:"items,omitempty" name:"items" title:"Top"`
	Page	*Page	`json:"page,omitempty" name:"page" title:"分页"`
}

