package request

// PassContactApplyRequest 通过联系人申请请求
type PassContactApplyRequest struct {
	OwnerId   string `json:"owner_id"`
	ContactId string `json:"contact_id"`
}
