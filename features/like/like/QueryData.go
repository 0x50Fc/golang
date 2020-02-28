package like

type QueryData struct {
	Items	[]*Like	`json:"items,omitempty" name:"items" title:"赞"`
	Page	*QueryDataPage	`json:"page,omitempty" name:"page" title:"分页"`
}

