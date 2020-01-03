package app

type QueryData struct {
	Items	[]*Ver	`json:"items,omitempty" name:"items" title:"版本"`
	Page	*QueryDataPage	`json:"page,omitempty" name:"page" title:"分页"`
}

