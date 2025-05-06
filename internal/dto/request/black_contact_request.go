package request

// BlackApplyRequest 黑名单消息请求
type BlackContactRequest struct {
	OwnerId   string `json:"owner_id"`
	ContactId string `json:"contact_id"`
}
