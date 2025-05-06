package request

// SmsLoginRequest 短信登录请求
type SmsLoginRequest struct {
	Telephone string `json:"telephone"`
	SmsCode   string `json:"sms_code"`
}
