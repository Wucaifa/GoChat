package request

// BlackApplyRequest 黑名单申请请求
type BlackApplyRequest struct {
	OwnerId   string `json:"owner_id"`
	ContactId string `json:"contact_id"`
}
