package request

// CreateGroupRequest 创建群组请求
type CreateGroupRequest struct {
	OwnerId string `json:"owner_id"` // 群主id
	Name    string `json:"name"`     // 群名称
	Notice  string `json:"notice"`   // 群公告
	AddMode int8   `json:"add_mode"` // 群添加模式
	Avatar  string `json:"avatar"`   // 群头像
}
