package request

// DeleteContactRequest 删除联系人请求
type DeleteContactRequest struct {
	OwnerId   string `json:"owner_id"`
	ContactId string `json:"contact_id"`
}
