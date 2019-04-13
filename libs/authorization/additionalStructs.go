package authorization

// Authorization will be in charge of account work
type Authorization struct {
	Token    string
	UserName string
}

// UserInfo is a struct for saving user's info
type UserInfo struct {
	ID       string `json:"id"`
	NickName string `json:"nickName,omitempty"`
	Gender   int    `json:"gender"`
}
