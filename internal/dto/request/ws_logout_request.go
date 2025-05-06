package request

// WsLogoutRequest WebSocket注销请求
type WsLogoutRequest struct {
	OwnerId string `json:"owner_id"`
}
