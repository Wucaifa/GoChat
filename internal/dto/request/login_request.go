package request

// LoginRequest 登录请求
type LoginRequest struct {
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}
