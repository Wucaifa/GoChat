package respond

// NewContactListRespond 新联系人列表响应
type NewContactListRespond struct {
	ContactId     string `json:"contact_id"`
	ContactName   string `json:"contact_name"`
	ContactAvatar string `json:"contact_avatar"`
	Message       string `json:"message"`
}
