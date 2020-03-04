package wx

type AppLoginData struct {
	SessionKey	string	`json:"session_key" name:"session_key" title:"session_key"`
	User	*User	`json:"user,omitempty" name:"user" title:"用户"`
}

