package request

// UpdateUserInfoRequest 更新用户信息请求
type UpdateUserInfoRequest struct {
	Uuid      string `json:"uuid"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	Birthday  string `json:"birthday"`
	Signature string `json:"signature"`
	Avatar    string `json:"avatar"`
}
