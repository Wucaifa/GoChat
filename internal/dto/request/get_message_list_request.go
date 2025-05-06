package request

// GetMessageListRequest 获取消息列表请求
type GetMessageListRequest struct {
	UserOneId string `json:"user_one_id"`
	UserTwoId string `json:"user_two_id"`
}
