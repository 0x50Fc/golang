package urelation

type FansQueryData struct {
	Items	[]*Fans	`json:"items,omitempty" title:"粉丝"`
	Page	*QueryDataPage	`json:"page,omitempty" title:"分页"`
}

