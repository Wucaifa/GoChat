package contact_status_enum

const (
	NORMAL         = iota // 正常
	BE_BLACK              // 被拉黑
	BLACK                 // 拉黑
	BE_DELETE             // 被删除
	DELETE                // 删除
	SILENCE               // 禁言
	QUIT_GROUP            // 退出群聊
	KICK_OUT_GROUP        // 被踢出群聊
)
