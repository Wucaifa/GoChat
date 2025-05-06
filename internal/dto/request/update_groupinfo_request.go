package request

// UpdateGroupInfoRequest 更新群组信息请求
type UpdateGroupInfoRequest struct {
	OwnerId string `json:"owner_id"`
	Uuid    string `json:"uuid"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	AddMode int8   `json:"add_mode"`
	Notice  string `json:"notice"`
}
