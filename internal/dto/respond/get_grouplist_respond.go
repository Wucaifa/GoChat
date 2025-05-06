package respond

// GetGroupListRespond 获取群组列表响应
type GetGroupListRespond struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	OwnerId   string `json:"owner_id"`
	Status    int8   `json:"status"`
	IsDeleted bool   `json:"is_deleted"`
}
