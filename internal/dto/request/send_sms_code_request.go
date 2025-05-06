package request

// SendSmsCodeRequest 发送短信验证码请求
type SendSmsCodeRequest struct {
	Telephone string `json:"telephone"`
}
