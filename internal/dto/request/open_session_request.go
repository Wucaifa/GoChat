package request

// OpenSessionRequest 打开会话请求
type OpenSessionRequest struct {
	SendId    string `json:"send_id"`
	ReceiveId string `json:"receive_id"`
}
