package request

// OwnlistRequest 获取用户列表请求
type OwnlistRequest struct {
	OwnerId string `json:"owner_id"`
}
