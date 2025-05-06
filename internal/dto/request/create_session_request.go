package request

// CreateSessionRequest 创建会话请求
type CreateSessionRequest struct {
	SendId    string `json:"send_id"`
	ReceiveId string `json:"receive_id"`
}
