package comment

type QueryData struct {
	Items	[]*Comment	`json:"items,omitempty" name:"items" title:"评论"`
	Page	*Page	`json:"page,omitempty" name:"page" title:"分页"`
}

