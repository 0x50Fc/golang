package wx

type AppPhoneData struct {
	Phone	string	`json:"phone" name:"phone" title:"手机号"`
	Country	string	`json:"country" name:"country" title:"国家区号"`
	User	*User	`json:"user,omitempty" name:"user" title:"用户"`
}

