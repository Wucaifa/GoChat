package request

// DeleteGroupsRequest 删除群组请求
type DeleteGroupsRequest struct {
	UuidList []string `json:"uuid_list"`
}
