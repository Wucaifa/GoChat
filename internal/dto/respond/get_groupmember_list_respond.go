package respond

// GetGroupMemberListRespond 获取群组成员列表响应
type GetGroupMemberListRespond struct {
	UserId   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
