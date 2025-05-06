package request

// RegisterRequest 用户注册请求
type RegisterRequest struct {
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	SmsCode   string `json:"sms_code"`
}
