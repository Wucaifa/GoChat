package request

// LeaveGroupRequest 退出群组请求
type LeaveGroupRequest struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`
}
