package respond

// LoadMyGroupRespond 加载我的群组响应
type LoadMyGroupRespond struct {
	GroupId   string `json:"group_id"`
	GroupName string `json:"group_name"`
	Avatar    string `json:"avatar"`
}
