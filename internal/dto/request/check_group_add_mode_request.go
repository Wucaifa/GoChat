package request

// CheckGroupAddModeRequest 检查群添加模式请求
// 📌 常见的群添加模式可能有：
// 无需验证：任何人都可以直接加入。

// 需要群主/管理员同意：加群申请需审批。

// 禁止加群：不允许外部用户加入。
type CheckGroupAddModeRequest struct {
	GroupId string `json:"group_id"`
}
