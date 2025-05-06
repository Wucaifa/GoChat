package request

// GetCurContactListInChatRoomRequest 获取当前聊天室联系人列表请求
type GetCurContactListInChatRoomRequest struct {
	OwnerId   string `json:"owner_id"`
	ContactId string `json:"contact_id"`
}
