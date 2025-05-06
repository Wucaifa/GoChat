package request

/*
* 一个典型的 DTO（数据传输对象） 结构体，定义了前端请求时传过来的参数格式。
 */
type AbleUsersRequest struct {
	UuidList []string `json:"uuid_list"` // 用户唯一id列表
	IsAdmin  int8     `json:"is_admin"`  // 是否是管理员，0.不是，1.是
}
