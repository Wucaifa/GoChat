package request

// RemoveGroupMembersRequest 删除群成员请求
type RemoveGroupMembersRequest struct {
	GroupId  string   `json:"group_id"`
	OwnerId  string   `json:"owner_id"`
	UuidList []string `json:"uuid_list"`
}
