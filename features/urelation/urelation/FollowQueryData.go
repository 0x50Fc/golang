package urelation

type FollowQueryData struct {
	Items	[]*Follow	`json:"items,omitempty" title:"关系"`
	Page	*QueryDataPage	`json:"page,omitempty" title:"分页"`
}

