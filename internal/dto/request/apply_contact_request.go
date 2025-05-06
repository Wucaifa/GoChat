package request

// ApplyContactRequest 申请添加好友请求
type ApplyContactRequest struct {
	OwnerId   string `json:"owner_id"`   // 申请人id
	ContactId string `json:"contact_id"` // 被申请人id
	Message   string `json:"message"`    // 申请消息
}
