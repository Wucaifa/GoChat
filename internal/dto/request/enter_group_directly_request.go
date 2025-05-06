package request

// EnterGroupDirectlyRequest 直接进入群组请求
type EnterGroupDirectlyRequest struct {
	OwnerId   string `json:"owner_id"`
	ContactId string `json:"contact_id"`
}
