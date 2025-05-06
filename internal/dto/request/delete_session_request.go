package request

// DeleteSessionRequest 删除会话请求
type DeleteSessionRequest struct {
	OwnerId   string `json:"owner_id"`
	SessionId string `json:"session_id"`
}
