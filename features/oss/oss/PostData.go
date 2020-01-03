package oss

type PostData struct {
	Url	string	`json:"url" name:"url" title:"上传URL"`
	Data	interface{}	`json:"data,omitempty" name:"data" title:"上传 Form Data"`
}

