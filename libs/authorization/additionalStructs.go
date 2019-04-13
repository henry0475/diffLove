package authorization

// Authorization will be in charge of account work
type Authorization struct {
	UserName string
}

// UserInfo is a struct for saving user's info
type UserInfo struct {
	ID       uint64 `json:"id"`
	UserName string `json:"username"`
	NickName string `json:"nickName,omitempty"`
	Gender   int    `json:"gender"`
	Token    string `json:"token"`
}
