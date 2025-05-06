package respond

// MyUserListRespond 我的用户列表响应
type MyUserListRespond struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
}
