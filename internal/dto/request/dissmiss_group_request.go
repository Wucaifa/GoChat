package request

// DismissGroupRequest 解散群组请求
type DismissGroupRequest struct {
	OwnerId string `json:"owner_id"`
	GroupId string `json:"group_id"`
}
