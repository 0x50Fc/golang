package media

type QueryData struct {
	Items	[]*Media	`json:"items,omitempty" name:"items" title:"媒体"`
	Page	*QueryDataPage	`json:"page,omitempty" name:"page" title:"分页"`
}

