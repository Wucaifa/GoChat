package request

// SetGroupsStatusRequest 设置群组状态请求
type SetGroupsStatusRequest struct {
	UuidList []string `json:"uuid_list"`
	Status   int8     `json:"status"`
}
